package service

import (
	"devflow/model"
	"github.com/xanzy/go-gitlab"
)

type ProjectServiceInterface interface {
	FetchProjects(pageNumber, pageSize int) ([]*model.Project, error)
	FetchProjectsCount() (int, error)
	RemoveProject(id int) (int64, error)
	ModifyProject(project model.Project) (int64, error)
	SaveProject(project model.Project) (int64, error)
	FetchBranches(gitlabId int) ([]*gitlab.Branch, error)
	FetchBranchesDetails(gitlabId int, branch string) (*gitlab.Branch, error)
	BuildProjectV2(params model.BuildParams, projectID int) (int64, error)
	DeployProject(r model.ProjectDeploy) (interface{}, error)
	GetBuildDetails(projectId int) (interface{}, error)
	GetBuildDetailsCount(projectId int) (int, error)
	ModifyProjectBuildStatus(deploymentName, status string, jobId int) (int64, error)
	GetBuildStatus() ([]int, error)
	FetchProjectsTags(projectName, env string) (interface{}, error)
}

type VmServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Vm, error)
	Count() (int, error)
	Create(vm model.Vm) (int64, error)
	Update(vm model.Vm) (int64, error)
	Delete(id int) (int64, error)
	FetchVmPasswordById(id int) (string, error)
	FetchVmsByApplication(application string) (interface{}, error)
}

type ImageServiceInterface interface {
	Fetch(pageNumber, pageSize int) ([]*model.Image, error)
	FetchImagesCount() (int, error)
	SaveImage(image model.Image) (int64, error)
	RemoveImage(id int) (int64, error)
	ModifyImage(image model.Image) (int64, error)
}

type EnvServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Env, error)
	Count() (int, error)
	Create(env model.Env) (int64, error)
	Update(env model.Env) (int64, error)
	Delete(id int) (int64, error)
	GetNsByEnv(env string) (interface{}, error)
}

type TaskServiceInterface interface {
	List(pageNumber, pageSize int) ([]*model.Task, error)
	Count() (int, error)
	Create(task model.Task) (int64, error)
	Update(task model.Task) (int64, error)
	Delete(id int) (int64, error)
}
