package todo

import (
	"database/sql"

	"github.com/samhodg1993/toto-todo-cli/internal/service"

	"github.com/spf13/cobra"
)

var db *sql.DB
var TodoService *service.TodoService

// TodoCmd represents the todo command group
var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage todos",
	Long:  `Create, list, update, and delete todo items.`,
}

// SetDatabase sets the database connection for the todo commands
func SetDatabase(database *sql.DB) {
	db = database
	TodoService = service.NewTodoService(db)
}

// Keeping this here for now. Might want to revert back to sub commands later
// func init() {
// Add subcommands
// TodoCmd.AddCommand(AddCmd)
// TodoCmd.AddCommand(DeleteTodo)
// TodoCmd.AddCommand(DelTodo)
// TodoCmd.AddCommand(ToggleComplete)
// TodoCmd.AddCommand(ToggleComp)
// TodoCmd.AddCommand(EditTodo)
// TodoCmd.AddCommand(LsCmd)
// TodoCmd.AddCommand(GetCmd)
// TodoCmd.AddCommand(GetCmdLong)
// TodoCmd.AddCommand(LsCmdLong)
// TodoCmd.AddCommand(LslCmdLong)
// }
