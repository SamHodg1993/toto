package todo

import (
	"fmt"

	"strconv"

	"github.com/spf13/cobra"
)

var sql_delete_todos string = "DELETE FROM todos WHERE id = ?"

var deleteSelectedId int = 0

var DeleteTodo = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id using the -i flag",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		id := strconv.Itoa(deleteSelectedId)

		err := TodoService.DeleteTodo(id)
		if err != nil {
			fmt.Printf("There was an error deleting the todo from the database: %v.\n", err)
			return
		}

		fmt.Printf("Todo with id: %v deleted successfully.\n", id)
	},
}

var DelTodo = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id using the -i flag",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		id := strconv.Itoa(deleteSelectedId)

		err := TodoService.DeleteTodo(id)
		if err != nil {
			fmt.Printf("There was an error deleting the todo from the database: %v.\n", err)
			return
		}

		fmt.Printf("Todo with id: %v deleted successfully.\n", id)
	},
}

func init() {
	DeleteTodo.Flags().IntVarP(&deleteSelectedId, "Todo ID", "i", 0, "The target todo's ID")
	DelTodo.Flags().IntVarP(&deleteSelectedId, "Todo ID", "i", 0, "The target todo's ID")
}
