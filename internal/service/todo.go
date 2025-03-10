package service

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

// TodoService handles todo operations
type TodoService struct {
	db             *sql.DB
	projectService *ProjectService
}

// NewTodoService creates a new todo service
func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		db:             db,
		projectService: NewProjectService(db),
	}
}

// GetTodosForFilepath gets todos for the current directory's project
func (s *TodoService) GetTodosForFilepath() (*sql.Rows, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	row := s.db.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	err = row.Scan(&projectId)
	if err != nil {
		choice, err := s.projectService.HandleNoExistingProject()
		if err != nil {
			return nil, err
		}
		if choice == 2 {
			return s.GetTodosForFilepath()
		}
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project")
		projectId = 1
	}

	rows, err := s.db.Query("SELECT id, title, completed FROM todos WHERE project_id = ?", projectId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

// GetAllTodos_LONG gets all todos with detailed information
func (s *TodoService) GetAllTodos_LONG() (*sql.Rows, error) {
	rows, err := s.db.Query(
		`SELECT 
			id, 
			title, 
			description, 
			project_id, 
			created_at, 
			updated_at, 
			completed 	
		FROM todos 
		`)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

// GetTodosForFilepath_LONG gets detailed todos for the current directory's project
func (s *TodoService) GetTodosForFilepath_LONG() (*sql.Rows, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	row := s.db.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	err = row.Scan(&projectId)
	if err != nil {
		var cancel string
		fmt.Println(`There is currently no project for this filepath. 
			Would you like to 
			0 - Cancel 
			1 - Add to the global todo list? 
			OR 
			2 - Create a new project for this filepath?`)
		fmt.Scanf("%s", &cancel)
		if cancel == "1" {
			projectId = 1
			err = nil
		} else if cancel == "2" {
			s.projectService.AddNewProjectWithPrompt()
			return s.GetTodosForFilepath_LONG()
		} else {
			fmt.Println("Aborting.")
			return nil, fmt.Errorf("operation cancelled by user")
		}
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project")
		projectId = 1
	}

	rows, err := s.db.Query(
		`SELECT 
			id, 
			title, 
			description, 
			project_id, 
			created_at, 
			updated_at, 
			completed 	
		FROM todos 
		WHERE project_id = ?
		`, projectId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

// AddTodo adds a new todo
func (s *TodoService) AddTodo(title, description string, projectId int, createdAt, updatedAt time.Time) error {
	var relevantProject int = 1

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
		"INSERT INTO todos (title, description, created_at, updated_at, project_id) VALUES (?,?,?,?,?)",
		title, description, createdAt, updatedAt, relevantProject)
	if err != nil {
		return fmt.Errorf("error adding todo: %w", err)
	}

	return nil
}

// DeleteTodo deletes a todo by ID
func (s *TodoService) DeleteTodo(id string) error {
	result, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no todo found with ID %s", id)
	}

	return nil
}

// ToggleComplete toggles the completion status of a todo
func (s *TodoService) ToggleComplete(id string) (bool, error) {
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
	result, err := s.db.Exec("UPDATE todos SET completed = ? WHERE id = ?", newStatus, id)
	if err != nil {
		return false, fmt.Errorf("error updating todo: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, fmt.Errorf("todo was not updated")
	}

	return newStatus, nil
}

// UpdateTodo updates a todo's title and/or description
func (s *TodoService) UpdateTodo(id int, title, description string, titleProvided, descProvided bool) (string, error) {
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
		finalTitle = title
	}
	if descProvided {
		finalDesc = description
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
