package jira

import (
	"database/sql"

	"github.com/samhodg1993/toto-todo-cli/internal/service"

	"github.com/spf13/cobra"
)

var db *sql.DB
var JiraService *service.JiraService

// JiraCmd represents the jira command group
var JiraCmd = &cobra.Command{
	Use:   "jira",
	Short: "Manage jira tickets",
	Long:  `Pull, Push and Sync tickets/todo's between toto and Jira`,
}

// SetDatabase sets the database connection for the jira commands
func SetDatabase(database *sql.DB) {
	db = database
	JiraService = service.NewJiraService(db)
}
