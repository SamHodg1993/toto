package todo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ODGY8/toto/internal/utilities"
)

// AddTodo adds a new todo
func (s *Service) AddTodo(
	title string,
	description string,
	projectId int,
	createdAt,
	updatedAt time.Time,
	jiraTicketRowId int64,
) error {
	var relevantProject int = 1

	sanitisedTitle := utilities.SanitizeInput(title, "title")
	sanitisedDesc := utilities.SanitizeInput(description, "description")

	// If projectId specified, use it
	if projectId != 0 {
		relevantProject = projectId
	} else {
		// Try to get project for current directory
		projectID, err := s.projectService.GetProjectIdByFilepath()
		if err != nil {
			// Handle no existing project
			choice, err := s.projectService.HandleNoExistingProject()
			if err != nil {
				return err
			}

			switch choice {
			case 0:
				return fmt.Errorf("operation cancelled by user")
			case 1:
				relevantProject = 1 // Global project
			case 2:
				// Create new project
				s.projectService.HandleAddNewProject("", "")
				projectID, err := s.projectService.GetProjectIdByFilepath()
				if err != nil {
					return fmt.Errorf("error getting project ID: %w", err)
				}
				relevantProject = projectID
			}
		} else {
			relevantProject = projectID
		}
	}

	// Insert the todo
	_, err := s.db.Exec(
		"INSERT INTO todos (title, description, created_at, updated_at, project_id, jira_ticket_id) VALUES (?,?,?,?,?,?)",
		sanitisedTitle, sanitisedDesc, createdAt, updatedAt, relevantProject, jiraTicketRowId)
	if err != nil {
		return fmt.Errorf("error adding todo: %w", err)
	}

	return nil
}

// DeleteTodo deletes a todo by ID
func (s *Service) DeleteTodo(ids []int) {
	for _, id := range ids {
		result, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)

		if err != nil {
			fmt.Printf("Unable to delete todo with ID: %d. Skipping...\n", id)
			continue
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Printf("Unable to delete todo with ID: %d. Skipping...\n", id)
			continue
		}

		if rowsAffected == 0 {
			fmt.Printf("No todo found with ID: %d. Skipping...\n", id)
			continue
		}

		fmt.Printf("Successfully deleted todo with ID: %d.\n", id)
	}

	fmt.Println("Completed deleting todos")
}

// UpdateTodo updates a todo's title and/or description
func (s *Service) UpdateTodo(id int, title string, description string, titleProvided, descProvided bool) (string, error) {
	// Get todo's current data
	var (
		todoTitle       string
		todoDescription sql.NullString
		todoID          int
	)

	err := s.db.QueryRow(
		"SELECT id, title, description FROM todos WHERE id = ?",
		id).Scan(&todoID, &todoTitle, &todoDescription)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("no todo found with ID %d", id)
	}
	if err != nil {
		return "", fmt.Errorf("error querying todo: %w", err)
	}

	// Set default values from existing todo
	finalTitle := todoTitle
	finalDesc := ""
	if todoDescription.Valid {
		finalDesc = todoDescription.String
	}

	// Override with provided values
	if titleProvided {
		sanitisedTitle := utilities.SanitizeInput(title, "title")
		finalTitle = sanitisedTitle
	}
	if descProvided {
		sanitisedDesc := utilities.SanitizeInput(description, "description")
		finalDesc = sanitisedDesc
	}

	// Update todo
	_, err = s.db.Exec(
		"UPDATE todos SET title = ?, description = ?, updated_at = ? WHERE id = ?",
		finalTitle, finalDesc, time.Now(), id)

	if err != nil {
		return "", fmt.Errorf("error updating todo: %w", err)
	}

	// Generate response message
	var message string
	if titleProvided && descProvided {
		message = fmt.Sprintf("Updated both title and description for todo #%d", id)
	} else if titleProvided {
		message = fmt.Sprintf("Updated title for todo #%d", id)
	} else if descProvided {
		message = fmt.Sprintf("Updated description for todo #%d", id)
	} else {
		message = fmt.Sprintf("No changes made to todo #%d", id)
	}

	return message, nil
}
