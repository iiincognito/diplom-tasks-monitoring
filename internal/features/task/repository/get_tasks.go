package repository

import (
	"fmt"
	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
)

func (r *TaskRepository) GetTasks(search string) ([]domain.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler`
	var args []interface{}

	if search != "" {
		// Check if search is a date in format 02.01.2006
		date, err := parseSearchDate(search)
		if err == nil {
			// Search by specific date
			query += ` WHERE date = ?`
			args = append(args, date)
		} else {
			// Search in title and comment using LIKE
			query += ` WHERE title LIKE ? OR comment LIKE ?`
			likePattern := "%" + search + "%"
			args = append(args, likePattern, likePattern)
		}
	}

	query += ` ORDER BY date ASC LIMIT 50`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func parseSearchDate(search string) (string, error) {
	// Try to parse date in format 02.01.2006
	if len(search) == 10 && search[2] == '.' && search[5] == '.' {
		day := search[0:2]
		month := search[3:5]
		year := search[6:10]
		// Return in format 20060102
		return year + month + day, nil
	}
	return "", fmt.Errorf("not a date format")
}
