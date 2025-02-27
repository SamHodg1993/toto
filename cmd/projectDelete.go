package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var sql_delete_projects string = "DELETE FROM projects WHERE id = ?"

var deleteProject = &cobra.Command{
	Use:   "project-delete",
	Short: "Delete a project",
	Long:  "Delete a single project from the database by referencing the project id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		projectId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("Invalid project ID format: %v\n", err)
			return
		}

		DeleteProject(projectId)

		fmt.Printf("Project with id: %v deleted successfully.\n", id)
	},
}

var delProj = &cobra.Command{
	Use:   "proj-del",
	Short: "Delete a project",
	Long:  "Delete a single project from the database by referencing the project id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		projectId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("Invalid project ID format: %v\n", err)
			return
		}

		err = DeleteProject(projectId)
		if err != nil {
			fmt.Printf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteProject)
	rootCmd.AddCommand(delProj)
}
