package cmd

import (
	"fmt"

	"database/sql"

	"github.com/spf13/cobra"
)

var sql_select_single_todo string = "SELECT id, completed FROM todos WHERE id = ?"
var sql_toggle_complete string = "UPDATE todos SET completed = ? WHERE id = ?"

var toggleComplete = &cobra.Command{
	Use:   "toggle-complete",
	Short: "Toggle a todo's status between complete and pending.",
	Long:  "Toggle an exising todo's status between complete and pending.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		completed := false

		// Query the current status
		err := database.QueryRow(sql_select_single_todo, id).Scan(&id, &completed)
		if err == sql.ErrNoRows {
			fmt.Printf("There is no todo with the id of %s\n", id)
			return
		}
		if err != nil {
			fmt.Printf("Error querying todo: %v\n", err)
			return
		}

		// Toggle the status
		newStatus := !completed

		// Update the todo
		result, err := database.Exec(sql_toggle_complete, newStatus, id)
		if err != nil {
			fmt.Printf("There was an error updating the todo in the database: %v\n", err)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected > 0 {
			fmt.Printf("Todo with id: %s status updated successfully\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(toggleComplete)
}
