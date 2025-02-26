package cmd

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/spf13/cobra"
)

var sql_update_todos string = "UPDATE todos SET title = ?, description = ?, updated_at = ? WHERE id = ?"
var sql_select_single_todo_for_update string = "SELECT id, title, description FROM todos WHERE id = ?"
var (
	targetId       int
	newDescription string
	newTitle       string
)

var editTodo = &cobra.Command{
	Use:   "edit",
	Short: "Update a todo",
	Long:  "Update a todo's title or description",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if ID is provided
		if targetId == 0 {
			fmt.Println("-i is a required flag.")
			return
		}

		// Check if any update flags were provided
		titleFlagProvided := cmd.Flags().Changed("title")
		descFlagProvided := cmd.Flags().Changed("description")

		if !titleFlagProvided && !descFlagProvided {
			fmt.Println("No changes specified. Please provide at least one field to update.")
			return
		}

		// Get todo's current data from db
		var (
			todoTitle       string
			todoDescription sql.NullString
			todoId          int
		)

		err := Database.QueryRow(sql_select_single_todo_for_update, targetId).Scan(&todoId, &todoTitle, &todoDescription)
		if err == sql.ErrNoRows {
			fmt.Printf("There is no todo with the id of %d.\n", targetId)
			return
		}
		if err != nil {
			fmt.Printf("Error querying todo: %v.\n", err)
			return
		}

		// Set default values from existing todo
		finalTitle := todoTitle
		finalDesc := ""
		if todoDescription.Valid {
			finalDesc = todoDescription.String
		}

		// Override with provided flags
		if titleFlagProvided {
			finalTitle = newTitle
		}

		if descFlagProvided {
			finalDesc = newDescription
		}

		// Update todo
		_, err = Database.Exec(sql_update_todos, finalTitle, finalDesc, time.Now(), targetId)
		if err != nil {
			fmt.Printf("Error updating todo: %v\n", err)
			return
		}

		// Report what was updated
		if titleFlagProvided && descFlagProvided {
			fmt.Printf("Updated both title and description for todo #%d.\n", targetId)
		} else if titleFlagProvided {
			fmt.Printf("Updated title for todo #%d.\n", targetId)
		} else if descFlagProvided {
			fmt.Printf("Updated description for todo #%d.\n", targetId)
		}
	},
}

func init() {
	editTodo.Flags().IntVarP(&targetId, "id", "i", 0, "Todo id")
	editTodo.Flags().StringVarP(&newTitle, "title", "t", "", "Todo title")
	editTodo.Flags().StringVarP(&newDescription, "description", "d", "", "Todo description")

	editTodo.MarkFlagRequired("id")

	rootCmd.AddCommand(editTodo)
}
