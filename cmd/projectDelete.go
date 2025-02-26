package cmd

import (
	"fmt"

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

		_, err := Database.Exec(sql_delete_projects, id)
		if err != nil {
			fmt.Printf("There was an error deleting the project from the database: %v.\n", err)
			return
		}

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

		_, err := Database.Exec(sql_delete_projects, id)
		if err != nil {
			fmt.Printf("There was an error deleting the project from the database: %v.\n", err)
			return
		}

		fmt.Printf("Project with id: %v deleted successfully.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteProject)
	rootCmd.AddCommand(delProj)
}
