package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/samhodg1993/todo-cli/cmd"
)

func GetTodosForFilepath() (*sql.Rows, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	row := cmd.Database.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	err = row.Scan(&projectId)
	if err != nil {
		choice, err := HandleNoExistingProject()
		if err != nil {
			return nil, err
		}
		if choice == 2 {
			GetTodosForFilepath()
		}
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project")
		projectId = 1
	}

	rows, err := cmd.Database.Query("SELECT id, title, completed FROM todos WHERE project_id = ?", projectId)
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
	rows, err := cmd.Database.Query(
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

	row := cmd.Database.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

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

	rows, err := cmd.Database.Query(
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
