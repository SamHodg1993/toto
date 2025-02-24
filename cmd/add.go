package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var sql_insert_todo string = "INSERT INTO todos (title, description, created_at, updated_at) VALUES (?,?,?,?)"

var (
	todoTitle       string
	todoDescription string
	todoCreatedAt   string
	todoUpdatedAt   string
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

		// If flag has value, update the created at value
		if todoCreatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoCreatedAt)
			if err == nil {
				createdAt = parsed
			}
		}

		// If flag has value, update the created at value
		if todoUpdatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoUpdatedAt)
			if err == nil {
				updatedAt = parsed
			}
		}

		_, err := database.Exec(sql_insert_todo, todoTitle, todoDescription, createdAt, updatedAt)
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

	addCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(addCmd)
}
