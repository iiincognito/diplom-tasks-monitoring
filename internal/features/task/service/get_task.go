package service

import (
	"strconv"
)

type GetTaskResponse struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
	Error   string `json:"error,omitempty"`
}

func (s *TaskService) GetTask(id int64) (*GetTaskResponse, error) {
	task, err := s.taskRepository.GetByID(id)
	if err != nil {
		return &GetTaskResponse{Error: err.Error()}, nil
	}
	if task == nil {
		return &GetTaskResponse{Error: "Задача не найдена"}, nil
	}

	return &GetTaskResponse{
		ID:      strconv.FormatInt(task.ID, 10),
		Date:    task.Date,
		Title:   task.Title,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}, nil
}
