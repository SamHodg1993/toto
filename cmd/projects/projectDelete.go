package projects

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var sql_delete_projects string = "DELETE FROM projects WHERE id = ?"

var deleteSelectedProjectId int

var DeleteProject = &cobra.Command{
	Use:   "project-delete",
	Short: "Delete a project",
	Long:  "Delete a single project from the database by referencing the project id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		id := strconv.Itoa(deleteSelectedProjectId)

		projectId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("Invalid project ID format: %v\n", err)
			return
		}

		ProjectService.DeleteProject(projectId)

		fmt.Printf("Project with id: %v deleted successfully.\n", id)
	},
}

var DelProj = &cobra.Command{
	Use:   "proj-del",
	Short: "Delete a project",
	Long:  "Delete a single project from the database by referencing the project id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		id := strconv.Itoa(deleteSelectedProjectId)

		projectId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("Invalid project ID format: %v\n", err)
			return
		}

		err = ProjectService.DeleteProject(projectId)
		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Printf("Project with id: %v deleted successfully.\n", id)
	},
}

func init() {
	DeleteProject.Flags().IntVarP(&deleteSelectedProjectId, "id", "i", 0, "The target project's ID")
	DelProj.Flags().IntVarP(&deleteSelectedProjectId, "id", "i", 0, "The target project's ID")
}
