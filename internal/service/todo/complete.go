package todo

import (
	"database/sql"
	"fmt"
)

// ToggleComplete toggles the completion status of a todo
func (s *Service) ToggleComplete(ids []int) {
	var todoID string
	var completed bool

	// Query current status
	for _, id := range ids {
		err := s.db.QueryRow("SELECT id, completed FROM todos WHERE id = ?", id).Scan(&todoID, &completed)

		if err == sql.ErrNoRows {
			fmt.Printf("no todo found with ID %d. Skipping...\n", id)
			continue
		}
		if err != nil {
			fmt.Printf("error querying todo: %w. Skipping...\n", err)
			continue
		}

		// Toggle status
		newStatus := !completed

		// Update todo
		result, err := s.db.Exec("UPDATE todos SET completed = ?, completed_at = CASE WHEN ? = 1 THEN CURRENT_TIMESTAMP ELSE NULL END WHERE id = ?", newStatus, newStatus, id)
		if err != nil {
			fmt.Printf("error updating todo: %w. Skipping...\n", err)
			continue
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			fmt.Printf("todo with ID %d was not updated.\n", id)
			continue
		}

		fmt.Printf("todo with ID %d successfully updated!\n", id)
	}
	fmt.Println("Completed todo toggling!")
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
