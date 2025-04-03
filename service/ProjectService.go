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
	"github.com/xanzy/go-gitlab"
	"io"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kubeErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type ProjectService struct {
	Repo             *repository.ProjectRepository
	ProjectBuildRepo *repository.ProjectBuildRepository
	TaskRepo         *repository.TaskRepository
	ImageRepo        *repository.ImageRepository
}

/*
FetchProjects 获取项目逻辑
*/

func (s *ProjectService) FetchProjects(pageNumber, pageSize int) (interface{}, error) {
	return s.Repo.GetProjects(pageNumber, pageSize)
}

func (s *ProjectService) FetchProjectsCount() (int, error) {
	return s.Repo.GetProjectsCount()
}

/*
RemoveProject 删除项目逻辑
*/

func (s *ProjectService) RemoveProject(id int) (int64, error) {
	return s.Repo.DeleteProject(id)
}

/*
ModifyProject 修改项目逻辑
*/

func (s *ProjectService) ModifyProject(project model.Project) (int64, error) {
	return s.Repo.UpdateProject(project)
}

/*
SaveProject 保存项目逻辑
*/

func (s *ProjectService) SaveProject(project model.Project) (int64, error) {
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

	if s.Repo.ExistDeploymentName(project.DeploymentName) {
		return 0, errors.New("请勿重复添加")
	}

	return s.Repo.CreateProject(project)
}

/*
FetchBranches 获取项目分支逻辑
*/

func (s *ProjectService) FetchBranches(gitlabId int) ([]*gitlab.Branch, error) {
	branches, _, err := GitlabClient.Branches.ListBranches(gitlabId, &gitlab.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (s *ProjectService) FetchBranchesDetails(gitlabId int, branch string) (interface{}, error) {
	b, _, err := GitlabClient.Branches.GetBranch(gitlabId, branch)
	if err != nil {
		return nil, err
	}
	return b, nil

}

/*
BuildProject 构建项目逻辑
*/

func (s *ProjectService) BuildProjectV2(params model.BuildParams, projectID int) (int64, error) {
	// 模版名称
	taskId, err := strconv.Atoi(params.TaskID)
	if err != nil {
		return 0, err
	}
	taskName, imageId, err := s.TaskRepo.GetTaskNameANDImageIDById(taskId)
	if err != nil {
		return 0, err
	}

	// 镜像名称
	imageName, err := s.ImageRepo.GetImageNameById(imageId)
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
	return s.ProjectBuildRepo.CreateProjectBuild(string(paramsJson), params.CreatedBy, taskName, projectID, jobBuild[0].Number+1)
}

//func (s *ProjectService) BuildProject(params model.BuildParams, createBy string, projectID int) (int64, error) {
//	buildTemplateID, err := s.Repo.GetBuildTemplateIDByID(projectID)
//	if err != nil {
//		return 0, err
//	}
//	buildTemplateName, _, err := s.BuildTemplateRepo.GetNameByID(buildTemplateID)
//	if err != nil {
//		return 0, err
//	}
//
//	paramsJson, err := json.Marshal(params)
//	if err != nil {
//		return 0, err
//	}
//
//	var paramsMap map[string]string
//	err = json.Unmarshal(paramsJson, &paramsMap)
//	if err != nil {
//		return 0, err
//	}
//
//	job, err := JenkinsClient.BuildJob(context.Background(), buildTemplateName, paramsMap)
//	if err != nil {
//		return 0, err
//	}
//	return s.ProjectBuildRepo.CreateProjectBuild(string(paramsJson), createBy, projectID, job+1)
//}

/*
DeployProject 发布项目逻辑
*/

func (s *ProjectService) DeployProject(r model.ProjectDeploy) (interface{}, error) {
	image := config.GlobalConfig.Harbor.URL + "/" + r.Env + "/" + r.Name + ":" + r.Tag

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
	return nil, errors.New("选择合适的发布方式")
}

/*
GetBuildDetails 获取项目构建详情逻辑
*/

func (s *ProjectService) GetBuildDetails(projectId int) (interface{}, error) {
	return s.ProjectBuildRepo.GetProjectBuildByProjectId(projectId)
}

func (s *ProjectService) GetBuildDetailsCount(projectId int) (interface{}, error) {
	return s.ProjectBuildRepo.GetProjectBuildCountByProjectId(projectId)
}

/*
ModifyProjectBuildStatus 修改项目构建状态逻辑
*/

func (s *ProjectService) ModifyProjectBuildStatus(deploymentName, status string, jobId int) (int64, error) {
	projectId, err := s.Repo.GetIdByDeploymentName(deploymentName)
	if err != nil {
		return 0, err
	}

	return s.ProjectBuildRepo.UpdateBuildStatus(projectId, jobId, status)
}

/*
GetBuildStatus 获取构建状态逻辑
*/

func (s *ProjectService) GetBuildStatus() (interface{}, error) {
	return s.ProjectBuildRepo.GetProjectIdByStatus()
}

/*
FetchProjectsTags 获取所有项目版本逻辑
*/

func (s *ProjectService) FetchProjectsTags(projectName, env string) (interface{}, error) {
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
