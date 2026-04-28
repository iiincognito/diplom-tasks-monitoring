package main

import (
	"context"
	dbConn "github.com/iiincognito/diplom-tasks-monitoring/internal/core/db"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/transport/http/server"
	task_http "github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/transport/http"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	_, err := dbConn.Init()
	if err != nil {
		log.Fatal(err)
	}
	tmp := ""
	taskTransport := task_http.NewTaskHTTPHandler(tmp)

	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	serv := server.NewHTTPServer(config)

	serv.Register(taskTransport.Register()...)

	if err := serv.Run(ctx); err != nil {
		log.Fatal(err)
	}

}
