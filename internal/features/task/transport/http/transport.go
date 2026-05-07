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

func (h *TaskHTTPHandler) Register(authMiddleware func(http.HandlerFunc) http.HandlerFunc) []server.Route {
	// Protected routes (require authentication if password is set)
	protectedRoutes := []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/task",
			Handler: authMiddleware(h.CreateTask),
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/task",
			Handler: authMiddleware(h.GetTask),
		},
		{
			Method:  http.MethodPut,
			Path:    "/api/task",
			Handler: authMiddleware(h.UpdateTask),
		},
		{
			Method:  http.MethodDelete,
			Path:    "/api/task",
			Handler: authMiddleware(h.DeleteTask),
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/task/done",
			Handler: authMiddleware(h.DoneTask),
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/tasks",
			Handler: authMiddleware(h.GetTasks),
		},
	}

	// Public routes (no authentication required)
	publicRoutes := []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/api/nextdate",
			Handler: h.NextDate,
		},
	}

	return append(protectedRoutes, publicRoutes...)
}
