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

func (e *EnvService) FetchEnvs(pageNumber, pageSize int) ([]*model.Env, error) {
	return e.EnvRepository.GetEnvs(pageNumber, pageSize)
}

func (e *EnvService) FetchEnvsCount() (int, error) {
	return e.EnvRepository.GetEnvsCount()
}

func (e *EnvService) FetchNamespacesByEnv(env string) (interface{}, error) {
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
