package task_http

import (
	"net/http"

	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/transport/http/server"
)

type TaskHTTPHandler struct {
	taskService TaskService
}

type TaskService interface {
	CreateTask(req *domain.CreateTaskRequest) (*domain.CreateTaskResponse, error)
}

func NewTaskHTTPHandler(taskService TaskService) *TaskHTTPHandler {
	return &TaskHTTPHandler{
		taskService: taskService,
	}
}

func (h *TaskHTTPHandler) Register() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/task",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/nextdate",
			Handler: h.NextDate,
		},
	}
}
