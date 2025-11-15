package todo

import (
	"database/sql"

	"github.com/odgy8/toto/internal/service/todo"

	"github.com/spf13/cobra"
)

var db *sql.DB
var TodoService *todo.Service

// TodoCmd represents the todo command group
var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage todos",
	Long:  `Create, list, update, and delete todo items.`,
}

// SetDatabase sets the database connection for the todo commands
func SetDatabase(database *sql.DB) {
	db = database
	TodoService = todo.New(db)
}
