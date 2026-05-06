package service

import (
	"fmt"
	"time"
)

type DoneTaskResponse struct {
	Error string `json:"error,omitempty"`
}

func (s *TaskService) DoneTask(id int64) (*DoneTaskResponse, error) {
	// Check if task exists
	existingTask, err := s.taskRepository.GetByID(id)
	if err != nil {
		return &DoneTaskResponse{Error: err.Error()}, nil
	}
	if existingTask == nil {
		return &DoneTaskResponse{Error: "Задача не найдена"}, nil
	}

	// If task is one-time (no repeat), delete it
	if existingTask.Repeat == "" {
		err = s.taskRepository.Delete(id)
		if err != nil {
			return &DoneTaskResponse{Error: err.Error()}, nil
		}
		return &DoneTaskResponse{}, nil
	}

	// For periodic tasks, calculate next date
	now := time.Now()
	nextDate, err := NextDate(now, existingTask.Date, existingTask.Repeat)
	if err != nil {
		return &DoneTaskResponse{Error: "Неверный формат правила повторения"}, nil
	}

	// Update task date
	existingTask.Date = nextDate
	err = s.taskRepository.Update(existingTask)
	if err != nil {
		fmt.Printf("DEBUG DoneTask Update error: %v\n", err)
		return &DoneTaskResponse{Error: err.Error()}, nil
	}

	fmt.Printf("DEBUG DoneTask success: id=%d, nextDate=%s\n", id, nextDate)
	return &DoneTaskResponse{}, nil
}
