package cmd

import (
	"database/sql"
	"fmt"
	"os"
)

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
		fmt.Printf("%v\n", err)
		return nil, err
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
		fmt.Printf("%v\n", err)
		return nil, err
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project")
		projectId = 1
	}

	rows, err := Database.Query("SELECT id, title, description, project_id, created_at, updated_at, completed FROM todos WHERE project_id = ?", projectId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}
