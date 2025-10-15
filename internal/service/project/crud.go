package project

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/samhodg1993/toto/internal/models"
	"github.com/samhodg1993/toto/internal/utilities"
)

// AddNewProject adds a new project to the database
func (s *Service) AddNewProject(project models.NewProject) error {
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

// AddNewProjectWithPrompt prompts the user for project details and adds the project
func (s *Service) AddNewProjectWithPrompt() error {
	var (
		project models.NewProject
		reader  = bufio.NewReader(os.Stdin)
	)

	fmt.Println("Please enter the title of your new project...")
	projectTitle, _ := reader.ReadString('\n')
	sanitisedTitle := utilities.SanitizeInput(projectTitle, "title")
	project.Title = strings.TrimSpace(sanitisedTitle)

	fmt.Println("Please enter the description of your new project...")
	projectDescription, _ := reader.ReadString('\n')
	sanitisedDesc := utilities.SanitizeInput(projectDescription, "description")
	project.Description = strings.TrimSpace(sanitisedDesc)

	return s.AddNewProject(project)
}

// HandleAddNewProject takes title and description from CLI and adds a project
func (s *Service) HandleAddNewProject(projectTitle string, projectDescription string) error {
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

	project.Title = utilities.SanitizeInput(project.Title, "title")
	project.Description = utilities.SanitizeInput(project.Description, "description")

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

// DeleteProject deletes a project and its associated todos
func (s *Service) DeleteProject(id int) error {
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

// UpdateProject updates a project's title, description, or filepath
func (s *Service) UpdateProject(projectID int, title string, description, filepath string, titleProvided, descProvided, filepathProvided bool) (string, error) {
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
		sanitisedTitle := utilities.SanitizeInput(title, "title")
		finalTitle = sanitisedTitle
	}
	if descProvided {
		sanitisedDesc := utilities.SanitizeInput(description, "description")
		finalDesc = sanitisedDesc
	}
	if filepathProvided {
		// Don't sanitize filepath - it needs exact formatting for paths
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
