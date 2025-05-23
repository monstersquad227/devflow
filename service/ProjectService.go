package service

import (
	"context"
	"devflow/config"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	pb "github.com/monstersquad227/flowedge-proto"
	"github.com/xanzy/go-gitlab"
	"io"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kubeErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type ProjectService struct {
	Repo             *repository.ProjectRepository
	ProjectBuildRepo *repository.ProjectBuildRepository
	TaskRepo         *repository.TaskRepository
	ImageRepo        *repository.ImageRepository
	FlowEdge         *FlowEdgeServer
}

func (svc *ProjectService) List(pageNumber, pageSize int) ([]*model.Project, error) {
	return svc.Repo.ListProjects(pageNumber, pageSize)
}

func (svc *ProjectService) Count() (int, error) {
	return svc.Repo.CountProjects()
}

func (svc *ProjectService) Create(project *model.Project) (int64, error) {
	projects, _, err := GitlabClient.Search.Projects(project.GitlabName, &gitlab.SearchOptions{})
	if err != nil {
		return 0, err
	}
	if len(projects) == 0 {
		return 0, errors.New("项目不存在")
	}
	for _, data := range projects {
		if data.Name == project.GitlabName {
			project.GitlabID = data.ID
			project.GitlabRepo = data.SSHURLToRepo
		}
	}

	if svc.Repo.ExistDeploymentName(project.DeploymentName) {
		return 0, errors.New("请勿重复添加")
	}

	return svc.Repo.CreateProject(project)
}

func (svc *ProjectService) Update(project *model.Project) (int64, error) {
	return svc.Repo.UpdateProject(project)
}

func (svc *ProjectService) Delete(id int) (int64, error) {
	return svc.Repo.DeleteProject(id)
}
func (svc *ProjectService) ListProjectApplications() ([]*model.Project, error) {
	return svc.Repo.ListProjectApplications()
}

func (svc *ProjectService) ListBranches(gitlabId int) ([]*gitlab.Branch, error) {
	branches, _, err := GitlabClient.Branches.ListBranches(gitlabId, &gitlab.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (svc *ProjectService) ListBranchesDetails(gitlabId int, branch string) (*gitlab.Branch, error) {
	b, _, err := GitlabClient.Branches.GetBranch(gitlabId, branch)
	if err != nil {
		return nil, err
	}
	return b, nil

}

func (svc *ProjectService) Build(params *model.BuildParams, projectID int) (int64, error) {
	// 模版名称
	taskId, err := strconv.Atoi(params.TaskID)
	if err != nil {
		return 0, err
	}
	taskName, imageId, err := svc.TaskRepo.GetTaskNameANDImageIDById(taskId)
	if err != nil {
		return 0, err
	}

	// 镜像名称
	imageName, err := svc.ImageRepo.GetImageNameById(imageId)
	if err != nil {
		return 0, err
	}
	if imageName == "" {
		return 0, errors.New("镜像不存在")
	}

	// 数据处理
	params.ImageSource = imageName
	paramsJson, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}
	var paramsMap map[string]string
	err = json.Unmarshal(paramsJson, &paramsMap)
	if err != nil {
		return 0, err
	}

	// 构建项目
	_, err = JenkinsClient.BuildJob(context.Background(), taskName, paramsMap)
	if err != nil {
		return 0, err
	}

	jobBuild, err := JenkinsClient.GetAllBuildIds(context.Background(), taskName)
	if err != nil {
		return 0, err
	}

	// 保存构建记录
	return svc.ProjectBuildRepo.CreateProjectBuild(string(paramsJson), params.CreatedBy, taskName, projectID, jobBuild[0].Number+1)
}

func (svc *ProjectService) Deploy(r *model.ProjectDeploy) (interface{}, error) {
	image := utils.GenerateImage(r.Env, r.Name, r.Tag)

	log.Printf(image)

	if r.PublishType == "kubernetes" {
		kubeClient, err := utils.KubernetesClient(r.Env + "config")
		if err != nil {
			return nil, err
		}
		deploy, err := kubeClient.AppsV1().Deployments(r.Namespace).Get(context.TODO(), r.Name, metav1.GetOptions{})

		if kubeErrors.IsNotFound(err) {
			return kubeClient.AppsV1().Deployments(r.Namespace).Create(
				context.TODO(),
				&appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:      r.Name,
						Namespace: r.Namespace,
					},
					Spec: appsv1.DeploymentSpec{
						Replicas: pointer.Int32(1),
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{"app": r.Name},
						},
						Template: corev1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{"app": r.Name},
							},
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name:  r.Name,
										Image: image,
									},
								},
							},
						},
					},
				},
				metav1.CreateOptions{})
		}

		if err != nil {
			return nil, err
		}

		deploy.Spec.Template.Spec.Containers[0].Image = image
		return kubeClient.AppsV1().Deployments(r.Namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	}

	if r.PublishType == "docker" {
		var wg sync.WaitGroup
		wg.Add(len(r.Ecs))
		for _, ecs := range r.Ecs {
			go func(addr string) {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						log.Printf("IP %s 出现错误: %v", addr, r)
					}
				}()
				// 删除容器
				utils.DeleteContainer(ecs, r.Name)
				// 下载镜像
				utils.PullImage(ecs, image)
				// 创建容器
				utils.CreateContainer(ecs, r.Name, image)
				// 启动容器
				utils.StartContainer(ecs, r.Name)
			}(ecs)
		}
		wg.Wait()
		return "success", nil
	}

	if r.PublishType == "flowedge" {
		responses := make(map[string]*pb.ExecuteResponse)
		for _, ecs := range r.Ecs {
			streamVal, ok := GlobalFlowEdgeServer.streams.Load(ecs)
			if !ok {
				return nil, errors.New("flowedge 节点 " + ecs + " 异常或未连接")
			}

			stream := streamVal.(pb.FlowEdge_CommunicateServer)
			commandID := uuid.New().String()
			respCh := make(chan *pb.ExecuteResponse, 1)
			GlobalFlowEdgeServer.pendingResponse.Store(commandID, respCh)
			err := stream.Send(&pb.StreamMessage{
				Type: pb.MessageType_EXECUTE_REQUEST,
				Body: &pb.StreamMessage_ExecuteRequest{
					ExecuteRequest: &pb.ExecuteRequest{
						CommandId:   commandID,
						Command:     "containerDragon",
						Image:       image,
						ContainerId: "",
					},
				},
			})
			if err != nil {
				GlobalFlowEdgeServer.pendingResponse.Delete(commandID)
				return nil, errors.New("发送到节点 " + ecs + " 失败: " + err.Error())
			}
			select {
			case resp := <-respCh:
				GlobalFlowEdgeServer.pendingResponse.Delete(commandID)
				responses[ecs] = resp
				// 可选：你可以将响应结果写入某个结果 map 或列表中统一返回
				log.Printf("节点 %s 响应: Output: %s, Error: %s", ecs, resp.Output, resp.Error)

			case <-time.After(60 * time.Second):
				GlobalFlowEdgeServer.pendingResponse.Delete(commandID)
				return nil, errors.New("等待节点 " + ecs + " 响应超时")
			}
		}
		return responses, nil
	}
	return nil, errors.New("选择合适的发布方式")
}

func (svc *ProjectService) ListBuildDetails(projectId int) (interface{}, error) {
	return svc.ProjectBuildRepo.GetProjectBuildByProjectId(projectId)
}

func (svc *ProjectService) ListBuildDetailsText(id int) (string, error) {
	jenkinsID, taskName, err := svc.ProjectBuildRepo.GetProjectBuildJenkinsIdAndTaskNameByProjectId(id)
	if err != nil {
		return "", err
	}

	url := "http://" + config.GlobalConfig.Jenkins.Username + ":" + config.GlobalConfig.Jenkins.Token + "@jenkins.chengdd.cn/job/" + taskName + "/" + jenkinsID + "/consoleText"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (svc *ProjectService) CountBuildDetails(projectId int) (int, error) {
	return svc.ProjectBuildRepo.GetProjectBuildCountByProjectId(projectId)
}

func (svc *ProjectService) ListBuildStatusING() ([]int, error) {
	return svc.ProjectBuildRepo.GetProjectIdByStatusING()
}

func (svc *ProjectService) ListBuildStatusFail() ([]int, error) {
	return svc.ProjectBuildRepo.GetProjectIdByStatusFail()
}

func (svc *ProjectService) UpdateBuildStatus(deploymentName, status string, jobId int) (int64, error) {
	projectId, err := svc.Repo.GetIdByDeploymentName(deploymentName)
	if err != nil {
		return 0, err
	}

	return svc.ProjectBuildRepo.UpdateBuildStatus(projectId, jobId, status)
}

func (svc *ProjectService) ListProjectImageTags(projectName, env string) (interface{}, error) {
	url := config.GlobalConfig.Harbor.URL + "/api/repositories/" + env + "%2F" + projectName + "/tags?detail=false"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	encrypt := base64.StdEncoding.EncodeToString([]byte(config.GlobalConfig.Harbor.Username + ":" + config.GlobalConfig.Harbor.Password))
	req.Header.Set("Authorization", "Basic "+encrypt)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	type harborTag struct {
		Name string `json:"name"`
	}
	var harborTagArray []*harborTag

	err = json.Unmarshal(bodyByte, &harborTagArray)
	if err != nil {
		return nil, err
	}
	return harborTagArray, nil
}
