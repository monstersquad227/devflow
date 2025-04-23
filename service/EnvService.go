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
	return e.EnvRepository.Get(pageNumber, pageSize)
}

func (e *EnvService) Fetch(pageNumber, pageSize int) ([]*model.Env, error) {
	return e.EnvRepository.Get(pageNumber, pageSize)
}

func (e *EnvService) FetchEnvsCount() (int, error) {
	return e.EnvRepository.GetEnvsCount()
}

func (e *EnvService) SaveEnv(env model.Env) (int64, error) {
	return e.EnvRepository.CreateEnv(env)
}

func (e *EnvService) RemoveEnv(id int) (int64, error) {
	return e.EnvRepository.DeleteEnv(id)
}

func (e *EnvService) ModifyEnv(env model.Env) (int64, error) {
	return e.EnvRepository.UpdateEnv(env)
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
