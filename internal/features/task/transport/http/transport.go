package task_http

import (
	"github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/service"
	"net/http"

	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/transport/http/server"
)

type TaskHTTPHandler struct {
	taskService TaskService
}

type TaskService interface {
	CreateTask(req *domain.CreateTaskRequest) (*domain.CreateTaskResponse, error)
	GetTasks(search string) (*service.GetTasksResponse, error)
	GetTask(id int64) (*service.GetTaskResponse, error)
	UpdateTask(req *service.UpdateTaskRequest) (*service.UpdateTaskResponse, error)
	DoneTask(id int64) (*service.DoneTaskResponse, error)
	DeleteTask(id int64) (*service.DeleteTaskResponse, error)
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
			Path:    "/api/task",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodPut,
			Path:    "/api/task",
			Handler: h.UpdateTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/api/task",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/task/done",
			Handler: h.DoneTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/nextdate",
			Handler: h.NextDate,
		},
	}
}
