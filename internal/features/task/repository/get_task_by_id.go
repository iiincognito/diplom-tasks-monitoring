package repository

import (
	"database/sql"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
)

func (r *TaskRepository) GetByID(id int64) (*domain.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	row := r.db.QueryRow(query, id)

	task := &domain.Task{}
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return task, nil
}
