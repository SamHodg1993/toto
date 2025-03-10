package service

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/samhodg1993/toto-todo-cli/internal/models"
)

var sql_insert_project string = `
	INSERT INTO projects (
		title, 
		description, 
		archived, 
		filepath, 
		created_at, 
		updated_at
	) VALUES (?,?,?,?,?,?)`

type ProjectService struct {
	db *sql.DB
}

func NewProjectService(db *sql.DB) *ProjectService {
	return &ProjectService{db: db}
}

// ListProjects returns all projects
func (s *ProjectService) ListProjects() (*sql.Rows, error) {
	return s.db.Query("SELECT id, title, filepath, archived FROM projects")
}

// GetProjectIdByFilepath returns the project ID for the current directory
func (s *ProjectService) GetProjectIdByFilepath() (int, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		return 0, fmt.Errorf("error getting current directory: %w", err)
	}

	row := s.db.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	err = row.Scan(&projectId)
	if err != nil {
		return 0, fmt.Errorf("no project exists for this filepath")
	}

	return projectId, nil
}

// AddNewProjectWithPrompt prompts the user for project details and adds the project
func (s *ProjectService) AddNewProjectWithPrompt() error {
	var (
		project models.NewProject
		reader  = bufio.NewReader(os.Stdin)
	)

	fmt.Println("Please enter the title of your new project...")
	projectTitle, _ := reader.ReadString('\n')
	project.Title = strings.TrimSpace(projectTitle)

	fmt.Println("Please enter the description of your new project...")
	projectDescription, _ := reader.ReadString('\n')
	project.Description = strings.TrimSpace(projectDescription)

	return s.AddNewProject(project)
}

// AddNewProject adds a new project to the database
func (s *ProjectService) AddNewProject(project models.NewProject) error {
	if strings.TrimSpace(project.Title) == "" {
		return fmt.Errorf("project title cannot be empty")
	}

	if strings.TrimSpace(project.Filepath) == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %w", err)
		}
		project.Filepath = currentDir
	}

	_, err := s.db.Exec(
		sql_insert_project,
		project.Title,
		project.Description,
		project.Archived,
		project.Filepath,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("error adding project: %w", err)
	}

	return nil
}

// DeleteProject deletes a project and its associated todos
func (s *ProjectService) DeleteProject(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid project id")
	}

	if id == 1 {
		return fmt.Errorf("cannot remove the global project as other functionality depends on it")
	}

	// Start a transaction to make sure no queries error out
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Delete todos
	_, err = tx.Exec("DELETE FROM todos WHERE project_id = ?", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting project todos: %w", err)
	}

	// Delete project
	result, err := tx.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting project: %w", err)
	}

	// Check if project existed
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("project with ID %d not found", id)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

// HandleNoExistingProject prompts the user for actions when no project exists
func (s *ProjectService) HandleNoExistingProject() (int, error) {
	var cancel string

	fmt.Println(`There is currently no project for this filepath. 
			Would you like to 
			0 - Cancel 
			1 - Add to the global todo list? 
			OR 
			2 - Create a new project for this filepath?`)
	fmt.Scanf("%s", &cancel)
	if cancel == "1" {
		return 1, nil
	} else if cancel == "2" {
		return 2, nil
	} else {
		fmt.Println("Aborting.")
		return 0, fmt.Errorf("operation cancelled by user")
	}
}

// UpdateProject updates a project's title, description, or filepath
func (s *ProjectService) UpdateProject(projectID int, title, description, filepath string, titleProvided, descProvided, filepathProvided bool) (string, error) {
	// Check if ID is valid
	if projectID <= 0 {
		return "", fmt.Errorf("invalid project ID")
	}

	// Get project's current data from db
	var (
		currentTitle       string
		currentDescription sql.NullString
		currentFilepath    string
	)

	err := s.db.QueryRow(
		"SELECT title, description, filepath FROM projects WHERE id = ?",
		projectID,
	).Scan(&currentTitle, &currentDescription, &currentFilepath)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("there is no project with the id of %d", projectID)
	}
	if err != nil {
		return "", fmt.Errorf("error querying project: %w", err)
	}

	// Set default values from existing project
	finalTitle := currentTitle
	finalDesc := ""
	finalFilepath := currentFilepath

	if currentDescription.Valid {
		finalDesc = currentDescription.String
	}

	// Override with provided values
	if titleProvided {
		finalTitle = title
	}
	if descProvided {
		finalDesc = description
	}
	if filepathProvided {
		finalFilepath = filepath
	}

	// Update project
	_, err = s.db.Exec(
		`UPDATE projects 
		 SET title = ?, description = ?, filepath = ?, updated_at = ? 
		 WHERE id = ?`,
		finalTitle, finalDesc, finalFilepath, time.Now(), projectID,
	)
	if err != nil {
		return "", fmt.Errorf("error updating project: %w", err)
	}

	// Generate response message
	var message string
	if titleProvided && descProvided && filepathProvided {
		message = fmt.Sprintf("Updated title, description and filepath for project #%d", projectID)
	} else if titleProvided && descProvided {
		message = fmt.Sprintf("Updated title and description for project #%d", projectID)
	} else if titleProvided && filepathProvided {
		message = fmt.Sprintf("Updated title and filepath for project #%d", projectID)
	} else if descProvided && filepathProvided {
		message = fmt.Sprintf("Updated description and filepath for project #%d", projectID)
	} else if titleProvided {
		message = fmt.Sprintf("Updated title for project #%d", projectID)
	} else if descProvided {
		message = fmt.Sprintf("Updated description for project #%d", projectID)
	} else if filepathProvided {
		message = fmt.Sprintf("Updated filepath for project #%d", projectID)
	} else {
		message = fmt.Sprintf("No changes made to project #%d", projectID)
	}

	return message, nil
}

// HandleAddNewProject takes title and description from CLI and adds a project
func (s *ProjectService) HandleAddNewProject(projectTitle string, projectDescription string) error {
	var project models.NewProject
	var reader = bufio.NewReader(os.Stdin)

	if projectTitle == "" {
		fmt.Println("Please enter the title of your new project...")
		input, _ := reader.ReadString('\n')
		project.Title = strings.TrimSpace(input)
	} else {
		project.Title = projectTitle
	}

	if projectDescription == "" {
		fmt.Println("Please enter the description of your new project...")
		input, _ := reader.ReadString('\n')
		project.Description = strings.TrimSpace(input)
	} else {
		project.Description = projectDescription
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}
	// Set the filepath
	project.Filepath = currentDir

	// Add the new project
	err = s.AddNewProject(project)
	if err != nil {
		return err
	}

	fmt.Printf("New project added: %s\n", project.Title)
	return nil
}
