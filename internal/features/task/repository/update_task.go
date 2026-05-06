package repository

import (
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
)

func (r *TaskRepository) Update(task *domain.Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	_, err := r.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	return err
}
