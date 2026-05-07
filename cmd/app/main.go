package main

import (
	"context"
	"net/http"

	dbConn "github.com/iiincognito/diplom-tasks-monitoring/internal/core/db"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/transport/http/server"
	auth_service "github.com/iiincognito/diplom-tasks-monitoring/internal/features/auth/service"
	auth_http "github.com/iiincognito/diplom-tasks-monitoring/internal/features/auth/transport/http"
	auth_middleware "github.com/iiincognito/diplom-tasks-monitoring/internal/features/auth/middleware"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/repository"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/service"
	task_http "github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/transport/http"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	db, err := dbConn.Init()
	if err != nil {
		log.Fatal(err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskSvc := service.NewTaskService(taskRepo)
	taskTransport := task_http.NewTaskHTTPHandler(taskSvc)

	authSvc := auth_service.NewAuthService()
	authMiddleware := auth_middleware.NewAuthMiddleware(authSvc)

	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	serv := server.NewHTTPServer(config)

	// Register public routes first
	serv.Register(server.Route{
		Method:  http.MethodPost,
		Path:    "/api/signin",
		Handler: auth_http.SignInHandler(authSvc),
	})

	// Register protected task routes with auth middleware
	taskRoutes := taskTransport.Register(authMiddleware)
	serv.Register(taskRoutes...)

	if err := serv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
