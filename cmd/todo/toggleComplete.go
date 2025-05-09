package todo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sql_select_single_todo string = "SELECT id, completed FROM todos WHERE id = ?"
var sql_toggle_complete string = "UPDATE todos SET completed = ? WHERE id = ?"

var ToggleComplete = &cobra.Command{
	Use:   "toggle-complete",
	Short: "Toggle a todo's status between complete and pending.",
	Long:  "Toggle an exising todo's status between complete and pending.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		newStatus, err := TodoService.ToggleComplete(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		statusText := "incomplete"
		if newStatus {
			statusText = "complete"
		}

		fmt.Printf("Todo with id: %s status updated successfully to %s\n", id, statusText)
	},
}

var ToggleComp = &cobra.Command{
	Use:   "comp",
	Short: "Toggle a todo's status between complete and pending.",
	Long:  "Toggle an exising todo's status between complete and pending.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		newStatus, err := TodoService.ToggleComplete(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		statusText := "incomplete"
		if newStatus {
			statusText = "complete"
		}

		fmt.Printf("Todo with id: %s status updated successfully to %s\n", id, statusText)
	},
}
