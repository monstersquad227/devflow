package service

import (
	"context"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvService struct {
	EnvRepository repository.EnvRepositoryInterface
}

// NewEnvService 构造函数，使用依赖注入
func NewEnvService(repo repository.EnvRepositoryInterface) *EnvService {
	return &EnvService{
		EnvRepository: repo,
	}
}

func (e *EnvService) List(pageNumber, pageSize int) ([]*model.Env, error) {
	return e.EnvRepository.ListEnvs(pageNumber, pageSize)
}

func (e *EnvService) Count() (int, error) {
	return e.EnvRepository.CountEnvs()
}

func (e *EnvService) Create(env *model.EnvCreateRequest) (*model.EnvCreateResponse, error) {
	LastInsertId, err := e.EnvRepository.CreateEnv(env)
	if err != nil {
		return nil, fmt.Errorf("创建 env 失败: %w", err)
	}
	resp := &model.EnvCreateResponse{
		LastInsertId: LastInsertId,
	}
	return resp, nil
}

func (e *EnvService) Update(env *model.Env) (int64, error) {
	return e.EnvRepository.UpdateEnv(env)
}

func (e *EnvService) Delete(id int) (int64, error) {
	return e.EnvRepository.DeleteEnv(id)
}

func (e *EnvService) GetNamespaces(env string) ([]v1.Namespace, error) {
	kubeClient, err := utils.KubernetesClient(env + "config")
	if err != nil {
		return nil, err
	}
	namespaces, err := kubeClient.CoreV1().Namespaces().List(
		context.Background(),
		metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces.Items, nil
}
