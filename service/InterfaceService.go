package service

import (
	"devflow/model"
	"github.com/xanzy/go-gitlab"
)

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
}

type ImageServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Image, error)
	Count() (int, error)
	Create(image *model.Image) (int64, error)
	Update(image *model.Image) (int64, error)
	Delete(id int) (int64, error)
}

type EnvServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Env, error)
	Count() (int, error)
	Create(env *model.Env) (int64, error)
	Update(env *model.Env) (int64, error)
	Delete(id int) (int64, error)
	GetNsByEnv(env string) (interface{}, error)
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
	Create(flowedge model.Flowedge) (int64, error)
}
