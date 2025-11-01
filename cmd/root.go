package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/samhodg1993/toto/cmd/jira"
	"github.com/samhodg1993/toto/cmd/projects"
	"github.com/samhodg1993/toto/cmd/todo"
	"github.com/samhodg1993/toto/cmd/utilityCommands"
	"github.com/samhodg1993/toto/internal/db"

	"github.com/spf13/cobra"
)

var Database *sql.DB

var RootCmd = &cobra.Command{
	Use:   "toto",
	Short: "A simple todo application for the command line.",
	Long:  "A simple CLI tool written to help organise my tasks and maintain a track of the important steps that I need to make. Initially this will be a command line only tool. But Jira integration would be kinda cool... In the future though.",
}

func Execute() {
	// Initialize database
	var err error
	Database, err = db.InitDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing database: %v\n", err)
		os.Exit(1)
	}
	defer Database.Close()

	// Pass database to command packages
	projects.SetDatabase(Database)
	todo.SetDatabase(Database)
	jira.SetDatabase(Database)
	utilityCommands.SetDatabase(Database)
	utilityCommands.InitDBService(Database)

	// Inject service dependencies
	// Jira service needs todo and project services
	jira.JiraService.SetDependencies(todo.TodoService, projects.ProjectService)
	// Todo service needs project service
	todo.TodoService.SetProjectService(projects.ProjectService)
	// Utility service needs todo and project services
	utilityCommands.UtilityService.SetDependencies(todo.TodoService, projects.ProjectService)

	// Add utility commands
	RootCmd.AddCommand(utilityCommands.ResetCmd)
	RootCmd.AddCommand(utilityCommands.CleanUtility)

	// Add the jira commands
	RootCmd.AddCommand(jira.JiraAuth)
	RootCmd.AddCommand(jira.JiraPull)
	RootCmd.AddCommand(jira.JiraPullClaude)

	// Add todo commands
	RootCmd.AddCommand(todo.AddCmd)
	RootCmd.AddCommand(todo.DeleteTodo)
	RootCmd.AddCommand(todo.DelTodo)
	RootCmd.AddCommand(todo.ToggleComplete)
	RootCmd.AddCommand(todo.ToggleComp)
	RootCmd.AddCommand(todo.EditTodo)
	RootCmd.AddCommand(todo.LsCmd)
	RootCmd.AddCommand(todo.GetCmd)
	RootCmd.AddCommand(todo.GetCmdLong)
	RootCmd.AddCommand(todo.LsCmdLong)
	RootCmd.AddCommand(todo.LslCmdLong)
	RootCmd.AddCommand(todo.RemoveCompleteForProject)
	RootCmd.AddCommand(todo.RemoveCompForProj)
	RootCmd.AddCommand(todo.GetTodoDescription)
	RootCmd.AddCommand(todo.GetTodoDesc)

	// Add project commands
	RootCmd.AddCommand(projects.ProjectAddCmd)
	RootCmd.AddCommand(projects.ProjAddCmd)
	RootCmd.AddCommand(projects.DeleteProject)
	RootCmd.AddCommand(projects.DelProj)
	RootCmd.AddCommand(projects.ProjectLsCmd)
	RootCmd.AddCommand(projects.ProjectListCmd)
	RootCmd.AddCommand(projects.EditProject)

	// Execute the root command
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
