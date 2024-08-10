package usecases

import (
	"task_manager/domains"
	"task_manager/repositories"
)

type TaskUsecase interface {
	GetTasks() ([]domains.Task, error)
	GetTaskByID(id string) (domains.Task, error)
	CreateTask(task domains.Task) error
	DeleteTask(id string) error
	UpdateTask(id string, task domains.Task) error
}

type taskUsecase struct {
	repo repositories.TaskRepository
}

func NewTaskUsecase(repo repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) GetTasks() ([]domains.Task, error) {
	return u.repo.GetAllTasks()
}

func (u *taskUsecase) GetTaskByID(id string) (domains.Task, error) {
	return u.repo.GetTaskByID(id)
}

func (u *taskUsecase) CreateTask(task domains.Task) error {
	return u.repo.CreateTask(task)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.repo.DeleteTask(id)
}

func (u *taskUsecase) UpdateTask(id string, task domains.Task) error {
	return u.repo.UpdateTask(id, task)
}
