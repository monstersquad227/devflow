package service

import (
	"devflow/model"
	"github.com/xanzy/go-gitlab"
	v1 "k8s.io/api/core/v1"
)

type UserServiceInterface interface {
	Login(account, password string) (interface{}, interface{}, error)
	List() ([]*model.User, error)
	UserPermission(account string) (interface{}, error)
}

type ProjectServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Project, error)
	Count() (int, error)
	Create(project *model.Project) (int64, error)
	Update(project *model.Project) (int64, error)
	Delete(id int) (int64, error)
	ListProjectApplications() ([]*model.Project, error)
	ListBranches(gitlabId int) ([]*gitlab.Branch, error)
	ListBranchesDetails(gitlabId int, branch string) (*gitlab.Branch, error)
	Build(params *model.BuildParams, projectID int) (int64, error)
	Deploy(r *model.ProjectDeploy) (interface{}, error)
	ListBuildDetails(projectId int) (interface{}, error)
	ListBuildDetailsText(id int) (string, error)
	CountBuildDetails(projectId int) (int, error)
	ListBuildStatusING() ([]int, error)
	ListBuildStatusFail() ([]int, error)
	UpdateBuildStatus(deploymentName, status string, jobId int) (int64, error)
	ListProjectImageTags(projectName, env string) (interface{}, error)
	ListProjectImageTagsV2(projectName, env string) (interface{}, error)
}

type VmServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Vm, error)
	Count() (int, error)
	Create(vm *model.Vm) (int64, error)
	Update(vm *model.Vm) (int64, error)
	Delete(id int) (int64, error)
	FetchVmPasswordById(id int) (string, error)
	FetchVmsByApplication(application string) (interface{}, error)
	CreateAliyunVm(vm *model.Vm) (int64, error)
	FetchUserByVm(id int) ([]*model.User, error)
	SetAssignUsersToVm(id int, users []int) (int64, error)
}

// ImageServiceInterface 定义镜像管理服务的接口
type ImageServiceInterface interface {
	// List 分页查询镜像列表
	List(pageNumber, pageSize int) ([]*model.Image, error)

	// Count 统计镜像总数
	Count() (int, error)

	// Create 创建新镜像
	Create(image *model.Image) (int64, error)

	// Update 更新镜像信息
	Update(image *model.Image) (int64, error)

	// Delete 删除指定ID的镜像
	Delete(id int) (int64, error)
}

// EnvServiceInterface 定义环境管理服务的接口
type EnvServiceInterface interface {
	// List 分页查询环境列表
	List(pageNumber, pageSize int) ([]*model.Env, error)

	// Count 统计环境总数
	Count() (int, error)

	// Create 创建新环境
	Create(env *model.EnvCreateRequest) (*model.EnvCreateResponse, error)

	// Update 更新环境信息
	Update(env *model.Env) (int64, error)

	// Delete 删除指定ID的环境
	Delete(id int) (int64, error)

	// GetNamespaces 获取指定环境下的所有命名空间
	GetNamespaces(env string) ([]v1.Namespace, error)
}

type TaskServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Task, error)
	Count() (int, error)
	Create(task *model.Task) (int64, error)
	Update(task *model.Task) (int64, error)
	Delete(id int) (int64, error)
}

type FlowedgeServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Flowedge, error)
	Count() (int, error)
	FetchFlowedgesByApplication(application string) (interface{}, error)
	Create(flowedge *model.Flowedge) (int64, error)
	UpdateApplication(flowedge *model.Flowedge) (int64, error)
}
