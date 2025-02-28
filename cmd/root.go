package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/samhodg1993/todo-cli/db"

	"github.com/spf13/cobra"
)

var Database *sql.DB

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo application for the command line.",
	Long:  "A simple CLI tool written to help organise my tasks and mainain a track of the important steps that I need to make. Initially this will be a command line only tool. But Jira integration would be kinda cool... In the future though. Right now I'm not even using Jira, I am however, using Monday.com... Hmmmm",
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

	// Execute the root command
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
