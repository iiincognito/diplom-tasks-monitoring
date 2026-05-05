package service

import (
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
)

type TaskService struct {
	taskRepository TaskRepository
}

type TaskRepository interface {
	Create(task *domain.Task) (int64, error)
	GetByID(id int64) (*domain.Task, error)
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{taskRepository: repo}
}
