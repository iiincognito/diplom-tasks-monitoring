package repository

import "github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"

func (r *TaskRepository) Create(task *domain.Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
