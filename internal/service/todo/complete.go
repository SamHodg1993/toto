package todo

import (
	"database/sql"
	"fmt"
)

// ToggleComplete toggles the completion status of a todo
func (s *Service) ToggleComplete(id string) (bool, error) {
	var todoID string
	var completed bool

	// Query current status
	err := s.db.QueryRow("SELECT id, completed FROM todos WHERE id = ?", id).Scan(&todoID, &completed)
	if err == sql.ErrNoRows {
		return false, fmt.Errorf("no todo found with ID %s", id)
	}
	if err != nil {
		return false, fmt.Errorf("error querying todo: %w", err)
	}

	// Toggle status
	newStatus := !completed

	// Update todo
	result, err := s.db.Exec("UPDATE todos SET completed = ?, completed_at = CASE WHEN ? = 1 THEN CURRENT_TIMESTAMP ELSE NULL END WHERE id = ?", newStatus, newStatus, id)
	if err != nil {
		return false, fmt.Errorf("error updating todo: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, fmt.Errorf("todo was not updated")
	}

	return newStatus, nil
}

// RemoveCompletedTodosForProject removes all completed todos for a given project
func (s *Service) RemoveCompletedTodosForProject(projectId int) error {
	res, err := s.db.Exec("delete from todos where project_id = ? and completed = 1", projectId)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Completed todo's removed for the current project.")

	return nil
}
