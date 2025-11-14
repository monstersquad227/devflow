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
	UpdateEnv(env *model.Env) (int64, error)
}
