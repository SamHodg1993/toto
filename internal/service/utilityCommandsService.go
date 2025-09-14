package service

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/samhodg1993/toto-todo-cli/internal/utilities"
)

type UtilityCommandsService struct {
	db             *sql.DB
	todoService    *TodoService
	projectService *ProjectService
}

// NewUtilityCommandsService creates a new utility commands service
func NewUtilityCommandsService(db *sql.DB) *UtilityCommandsService {
	return &UtilityCommandsService{
		db:             db,
		todoService:    NewTodoService(db),
		projectService: NewProjectService(db),
	}
}

func (s *UtilityCommandsService) CleanAndPrintTodos() error {
	utilities.ClearScreen()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	row := s.db.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	var projectId int = 0
	err = row.Scan(&projectId)
	if err != nil {
		choice, err := s.projectService.HandleNoExistingProject()
		if err != nil {
			return err
		}
		if choice == 2 {
			return s.CleanAndPrintTodos()
		}
	}

	if projectId == 0 {
		fmt.Printf("No project exists for this working directory, defaulting to the global project\n")
		projectId = 1
	}

	err = s.todoService.RemoveCompletedTodosForProject(projectId)
	if err != nil {
		return err
	}

	rows, err := s.db.Query("SELECT id, title, completed FROM todos WHERE project_id = ?", projectId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	defer rows.Close()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Todo", "Status"})
	table.SetBorder(true)
	table.SetRowLine(true)

	strikethrough := color.New(color.CrossedOut).SprintFunc()

	for rows.Next() {
		var (
			id        int
			title     string
			completed bool
		)

		err = rows.Scan(&id, &title, &completed)
		if err != nil {
			return err
		}

		status := "Pending"
		if completed {
			status = strikethrough("Completed")
			title = strikethrough(title)
		}

		table.Append([]string{strconv.Itoa(id), title, status})
	}

	if err = rows.Err(); err != nil {
		return err
	}

	table.Render()
	return nil
}
