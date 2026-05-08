package repository

import (
	dbConn "github.com/iiincognito/diplom-tasks-monitoring/internal/core/db"
)

type TaskRepository struct {
	db *dbConn.DB
}

func NewTaskRepository(db *dbConn.DB) *TaskRepository {
	return &TaskRepository{db: db}
}
