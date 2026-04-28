package task_http

import (
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/transport/http/server"
	"net/http"
)

type TaskHTTPHandler struct {
	taskService TaskService
}

type TaskService interface{}

func NewTaskHTTPHandler(taskService TaskService) *TaskHTTPHandler {
	return &TaskHTTPHandler{
		taskService: taskService,
	}
}

func (h *TaskHTTPHandler) Register() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/task",
			Handler: h.GetTask,
		},
	}
}
