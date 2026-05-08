package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
)

type UpdateTaskResponse struct {
	Error string `json:"error,omitempty"`
}

type UpdateTaskRequest struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (s *TaskService) UpdateTask(req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	// Validate ID
	if req.ID == "" {
		return &UpdateTaskResponse{Error: "Не указан идентификатор"}, nil
	}
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return &UpdateTaskResponse{Error: "Неверный идентификатор"}, nil
	}

	// Check if task exists
	existingTask, err := s.taskRepository.GetByID(id)
	if err != nil {
		return &UpdateTaskResponse{Error: err.Error()}, nil
	}
	if existingTask == nil {
		return &UpdateTaskResponse{Error: "Задача не найдена"}, nil
	}

	// Validate title
	if req.Title == "" {
		return &UpdateTaskResponse{Error: "Не указан заголовок задачи"}, nil
	}

	now := time.Now()
	today := now.Format("20060102")

	date := req.Date
	if date == "" {
		date = today
	} else {
		if _, err := time.Parse("20060102", date); err != nil {
			return &UpdateTaskResponse{Error: "Неверный формат даты"}, nil
		}
	}

	if req.Repeat != "" {
		if date < today {
			nextDate, err := NextDate(now, date, req.Repeat)
			if err != nil {
				return &UpdateTaskResponse{Error: "Неверный формат правила повторения"}, nil
			}
			date = nextDate
		}
	} else if date < today {
		date = today
	}

	task := &domain.Task{
		ID:      id,
		Date:    date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}

	err = s.taskRepository.Update(task)
	if err != nil {
		fmt.Printf("DEBUG UpdateTask error: %v\n", err)
		return &UpdateTaskResponse{Error: err.Error()}, nil
	}

	fmt.Printf("DEBUG UpdateTask success: id=%d\n", id)
	return &UpdateTaskResponse{}, nil
}
