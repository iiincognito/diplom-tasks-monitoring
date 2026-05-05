package service

import (
	"fmt"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
	"time"
)

func (s *TaskService) CreateTask(req *domain.CreateTaskRequest) (*domain.CreateTaskResponse, error) {
	if req.Title == "" {
		return &domain.CreateTaskResponse{Error: "Не указан заголовок задачи"}, nil
	}

	now := time.Now()
	today := now.Format("20060102")

	date := req.Date
	if date == "" {
		date = today
	} else {
		if _, err := time.Parse("20060102", date); err != nil {
			return &domain.CreateTaskResponse{Error: "Неверный формат даты"}, nil
		}
	}

	if req.Repeat != "" {
		if date < today {
			nextDate, err := NextDate(now, date, req.Repeat)
			if err != nil {
				return &domain.CreateTaskResponse{Error: "Неверный формат правила повторения"}, nil
			}
			date = nextDate
		}
	} else if date < today {
		date = today
	}

	task := &domain.Task{
		Date:    date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}

	id, err := s.taskRepository.Create(task)
	if err != nil {
		fmt.Printf("DEBUG CreateTask error: %v\n", err)
		return &domain.CreateTaskResponse{Error: err.Error()}, nil
	}

	fmt.Printf("DEBUG CreateTask success: id=%d, task=%+v\n", id, task)
	return &domain.CreateTaskResponse{ID: id}, nil
}
