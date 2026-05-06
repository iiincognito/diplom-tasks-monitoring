package repository

func (r *TaskRepository) Delete(id int64) error {
	query := `DELETE FROM scheduler WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
