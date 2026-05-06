package service

import (
	"fmt"
	"strconv"
)

type TaskResponse struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type GetTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
	Error string         `json:"error,omitempty"`
}

func (s *TaskService) GetTasks(search string) (*GetTasksResponse, error) {
	tasks, err := s.taskRepository.GetTasks(search)
	if err != nil {
		fmt.Printf("DEBUG GetTasks error: %v\n", err)
		return &GetTasksResponse{Error: err.Error()}, nil
	}

	taskResponses := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		taskResponses = append(taskResponses, TaskResponse{
			ID:      strconv.FormatInt(task.ID, 10),
			Date:    task.Date,
			Title:   task.Title,
			Comment: task.Comment,
			Repeat:  task.Repeat,
		})
	}

	return &GetTasksResponse{Tasks: taskResponses}, nil
}
