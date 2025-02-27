package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/samhodg1993/todo-cli/models"
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

func GetProjectIdByFilepath() (int, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return 0, err
	}

	row := Database.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	err = row.Scan(&projectId)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, err
	}

	return projectId, nil
}

func GetTodosForFilepath() (*sql.Rows, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	row := Database.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

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
			AddNewProject_WITH_PROMPT()
			return GetTodosForFilepath()
		} else {
			fmt.Println("Aborting.")
			return nil, fmt.Errorf("operation cancelled by user")
		}
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project")
		projectId = 1
	}

	rows, err := Database.Query("SELECT id, title, completed FROM todos WHERE project_id = ?", projectId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

func GetAllTodos_LONG() (*sql.Rows, error) {
	rows, err := Database.Query(
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

func GetTodosForFilepath_LONG() (*sql.Rows, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	row := Database.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

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
			AddNewProject_WITH_PROMPT()
			return GetTodosForFilepath_LONG()
		} else {
			fmt.Println("Aborting.")
			return nil, fmt.Errorf("operation cancelled by user")
		}
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project")
		projectId = 1
	}

	rows, err := Database.Query(
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

func AddNewProject_WITH_PROMPT() {
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

	AddNewProject(project)
}

func AddNewProject(project models.NewProject) {
	if strings.TrimSpace(project.Title) == "" {
		fmt.Println("Project title cannot be empty")
		return
	}

	if strings.TrimSpace(project.Filepath) == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
		project.Filepath = currentDir
	}

	_, err := Database.Exec(
		sql_insert_project,
		project.Title,
		project.Description,
		project.Archived,
		project.Filepath,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		fmt.Printf("There was an error adding the project: %v\n", err)
		return
	}
	fmt.Printf("New project added: %s.\n", project.Title)
}

func DeleteProject(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid project id")
	}

	if id == 1 {
		return fmt.Errorf("Please do not remove the global project. Other functionality relies upon it. A fix to this is in the roadmap, but for right now, please allow the global project to remain.\n")
	}

	// Start a transaction to make sure no queries error out
	tx, err := Database.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Delete todos
	result, err := tx.Exec("DELETE FROM todos WHERE project_id = ?", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting project todos: %v", err)
	}

	// Delete project
	result, err = tx.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting project: %v", err)
	}

	// Check if project existed
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("project with ID %d not found", id)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	fmt.Printf("Project with ID %d and all associated todos deleted successfully.\n", id)
	return nil
}
