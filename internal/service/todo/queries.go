package todo

import (
	"fmt"
	"os"

	"github.com/ODGY8/toto/internal/models"
)

// GetTodosForFilepath gets todos for the current directory's project
func (s *Service) GetTodosForFilepath() ([]models.Todo, error) {
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
			s.projectService.AddNewProjectWithPrompt()
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
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo, err := scanRowToSimpleTodo(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning todo: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

// GetAllTodosForFilepath gets all todos regardless of project
func (s *Service) GetAllTodosForFilepath() ([]models.Todo, error) {
	rows, err := s.db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo, err := scanRowToSimpleTodo(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning todo: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

// GetAllTodos_LONG gets all todos with detailed information
func (s *Service) GetAllTodos_LONG() ([]models.Todo, error) {
	rows, err := s.db.Query(
		`SELECT
			id,
			title,
			description,
			project_id,
			created_at,
			updated_at,
			completed,
		  completed_at
		FROM todos
		`)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo, err := scanRowToTodo(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning todo: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

// GetAllTodos gets all todos regardless of filepath
func (s *Service) GetAllTodos() ([]models.Todo, error) {
	rows, err := s.db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo, err := scanRowToSimpleTodo(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning todo: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

// GetTodosForFilepath_LONG gets detailed todos for the current directory's project
func (s *Service) GetTodosForFilepath_LONG() ([]models.Todo, error) {
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
		if choice == 1 {
			projectId = 1
		} else if choice == 2 {
			s.projectService.AddNewProjectWithPrompt()
			return s.GetTodosForFilepath_LONG()
		} else {
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
			completed,
		  completed_at
		FROM todos
		WHERE project_id = ?
		`, projectId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo, err := scanRowToTodo(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning todo: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

// GetTodoDetails gets the description for a todo by ID
func (s *Service) GetTodoDetails(todoId int) (string, error) {
	result := s.db.QueryRow("SELECT description FROM todos WHERE id = ?", todoId)

	description := ""
	err := result.Scan(&description)
	if err != nil {
		return "", err
	}

	return description, nil
}
