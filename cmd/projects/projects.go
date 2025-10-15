package projects

import (
	"database/sql"

	"github.com/samhodg1993/toto/internal/service"

	"github.com/spf13/cobra"
)

var db *sql.DB
var ProjectService *service.ProjectService

// ProjectsCmd represents the projects command group
var ProjectsCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  `Create, list, update, and delete projects.`,
}

// SetDatabase sets the database connection for the projects commands
func SetDatabase(database *sql.DB) {
	db = database
	ProjectService = service.NewProjectService(db)
}
