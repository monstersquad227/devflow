package service

import (
	"context"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvService struct {
	EnvRepository *repository.EnvRepository
}

func (e *EnvService) List(pageNumber, pageSize int) ([]*model.Env, error) {
	return e.EnvRepository.ListEnvs(pageNumber, pageSize)
}

func (e *EnvService) Count() (int, error) {
	return e.EnvRepository.CountEnvs()
}

func (e *EnvService) Create(env *model.Env) (int64, error) {
	return e.EnvRepository.CreateEnv(env)
}

func (e *EnvService) Update(env *model.Env) (int64, error) {
	return e.EnvRepository.UpdateEnv(env)
}

func (e *EnvService) Delete(id int) (int64, error) {
	return e.EnvRepository.DeleteEnv(id)
}

func (e *EnvService) GetNsByEnv(env string) (interface{}, error) {
	kubeClient, err := utils.KubernetesClient(env + "config")
	if err != nil {
		return nil, err
	}
	namespaces, err := kubeClient.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces.Items, nil
}
