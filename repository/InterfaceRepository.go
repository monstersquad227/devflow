package repository

import "devflow/model"

// EnvRepositoryInterface 环境数据访问层接口
type EnvRepositoryInterface interface {
	// ListEnvs 分页查询环境列表
	ListEnvs(pageNumber, pageSize int) ([]*model.Env, error)

	// CountEnvs 统计环境总数
	CountEnvs() (int, error)

	// CreateEnv 创建环境
	CreateEnv(env *model.EnvCreateRequest) (int64, error)

	// DeleteEnv 删除环境
	DeleteEnv(id int) (int64, error)

	// UpdateEnv 更新环境
	UpdateEnv(env *model.EnvUpdateRequest) (int64, error)
}

type ImageRepositoryInterface interface {
	// ListImages 分页查询镜像列表
	ListImages(pageNumber, pageSize int) ([]*model.Image, error)

	// CountImages 统计镜像总数
	CountImages() (int, error)

	// CreateImage 创建镜像
	CreateImage(image *model.Image) (int64, error)

	// UpdateImage 更新镜像
	UpdateImage(image *model.Image) (int64, error)

	// DeleteImage 删除镜像
	DeleteImage(id int) (int64, error)

	// GetImageName 通过镜像ID获取镜像名称
	GetImageName(id int) (string, error)
}

type TaskRepositoryInterface interface {
	ListTasks(pageNumber, pageSize int) ([]*model.Task, error)
	CountTasks() (int, error)
	CreateTask(task *model.Task) (int64, error)
	DeleteTask(id int) (int64, error)
	UpdateTask(task *model.Task) (int64, error)
	GetTaskNameANDImageIDById(id int) (string, int, error)
}

type FlowedgeRepositoryInterface interface {
	ListFlowedges(pageNumber, pageSize int) ([]*model.Flowedge, error)
	CountFlowedges() (int, error)
	GetFlowedgeByApplication(application string) (interface{}, error)
	CreateFlowedge(flowedge *model.Flowedge) (int64, error)
	UpdateFlowedgeLastHeartBeat(flow *model.Flowedge) (int64, error)
	UpdateFlowedgeApplication(flow *model.FlowedgePatchRequest) (int64, error)
}
