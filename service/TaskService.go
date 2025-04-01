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

func (t *TaskService) SaveTask(task model.Task) (int64, error) {
	return t.TaskRepository.CreateTask(task)
}

func (t *TaskService) ModifyTask(task model.Task) (int64, error) {
	return t.TaskRepository.UpdateTask(task)
}
func (t *TaskService) RemoveTask(id int) (int64, error) {
	return t.TaskRepository.DeleteTask(id)
}
