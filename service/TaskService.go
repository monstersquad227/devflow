package service

import (
	"devflow/model"
	"devflow/repository"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

func (t *TaskService) FetchTasks(pageNumber, pageSize int) ([]*model.Task, error) {
	return t.TaskRepository.GetTasks(pageNumber, pageSize)
}

func (t *TaskService) FetchTasksCount() (int, error) {
	return t.TaskRepository.GetTasksCount()
}
