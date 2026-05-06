package service

type DeleteTaskResponse struct {
	Error string `json:"error,omitempty"`
}

func (s *TaskService) DeleteTask(id int64) (*DeleteTaskResponse, error) {
	// Check if task exists
	existingTask, err := s.taskRepository.GetByID(id)
	if err != nil {
		return &DeleteTaskResponse{Error: err.Error()}, nil
	}
	if existingTask == nil {
		return &DeleteTaskResponse{Error: "Задача не найдена"}, nil
	}

	err = s.taskRepository.Delete(id)
	if err != nil {
		return &DeleteTaskResponse{Error: err.Error()}, nil
	}

	return &DeleteTaskResponse{}, nil
}
