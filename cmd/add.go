package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var sql_insert_todo string = "INSERT INTO todos (title, description, created_at, updated_at, project_id) VALUES (?,?,?,?,?)"

var (
	todoTitle       string
	todoDescription string
	todoCreatedAt   string
	todoUpdatedAt   string
	todoProjectId   int
)

var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a new todo",
	Long:  "Add a new todo to the list of stored todos",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// This has moved to the -t flag
		// todo := args[0]

		// Create default timestamps
		createdAt := time.Now()
		updatedAt := time.Now()
		releventProject := 1

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		fmt.Printf("Working dir: %s", currentDir)

		// If created at flag has value, update the created at value
		if todoCreatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoCreatedAt)
			if err == nil {
				createdAt = parsed
			}
		}

		// If project id flag has value, update the projectId
		if todoProjectId != 0 {
			releventProject = todoProjectId
		}

		// If updated at flag has value, update the created at value
		if todoUpdatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoUpdatedAt)
			if err == nil {
				updatedAt = parsed
			}
		}

		// Update the todo
		_, err = database.Exec(sql_insert_todo, todoTitle, todoDescription, createdAt, updatedAt, releventProject)
		if err != nil {
			fmt.Printf("There was an error adding the todo: %v\n", err)
			return
		}
		fmt.Printf("New todo added: %s\n", todoTitle)
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&todoTitle, "title", "t", "", "Title of the todo")
	addCmd.PersistentFlags().StringVarP(&todoDescription, "description", "d", "", "Description of the todo")
	addCmd.PersistentFlags().StringVarP(&todoCreatedAt, "created-at", "c", "", "Todo creation time")
	addCmd.PersistentFlags().StringVarP(&todoUpdatedAt, "updated-at", "u", "", "Todo last updated time")
	addCmd.PersistentFlags().IntVarP(&todoProjectId, "project-id", "p", 0, "Relevent Project Id")

	addCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(addCmd)
}
