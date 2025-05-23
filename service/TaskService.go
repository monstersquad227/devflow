package service

import (
	"devflow/model"
	"devflow/repository"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

func (t *TaskService) List(pageNumber, pageSize int) ([]*model.Task, error) {
	return t.TaskRepository.ListTasks(pageNumber, pageSize)
}

func (t *TaskService) Count() (int, error) {
	return t.TaskRepository.CountTasks()
}

func (t *TaskService) Create(task *model.Task) (int64, error) {
	return t.TaskRepository.CreateTask(task)
}

func (t *TaskService) Update(task *model.Task) (int64, error) {
	return t.TaskRepository.UpdateTask(task)
}

func (t *TaskService) Delete(id int) (int64, error) {
	return t.TaskRepository.DeleteTask(id)
}
