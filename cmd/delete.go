package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sql_delete_todos string = "DELETE FROM todos WHERE id = ?"

var deleteTodo = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		_, err := database.Exec(sql_delete_todos, id)
		if err != nil {
			fmt.Printf("There was an error deleting the todo from the database: %v.\n", err)
			return
		}

		fmt.Printf("Todo with id: %v deleted successfully.\n", id)
	},
}

var delTodo = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		_, err := database.Exec(sql_delete_todos, id)
		if err != nil {
			fmt.Printf("There was an error deleting the todo from the database: %v.\n", err)
			return
		}

		fmt.Printf("Todo with id: %v deleted successfully.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteTodo)
	rootCmd.AddCommand(delTodo)
}
