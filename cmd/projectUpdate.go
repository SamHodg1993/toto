package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var sql_update_project string = "UPDATE projects SET title = ?, description = ?, updated_at = ? WHERE id = ?"
var sql_select_single_project_for_update string = "SELECT id, title, description FROM projects WHERE id = ?"
var (
	targetProjectId       int
	newProjectDescription string
	newProjectTitle       string
)
var editProject = &cobra.Command{
	Use:   "project-edit",
	Short: "Update a project",
	Long:  "Update a project's title or description",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if ID is provided
		if targetProjectId == 0 {
			fmt.Println("-i is a required flag.")
			return
		}
		// Check if any update flags were provided
		titleFlagProvided := cmd.Flags().Changed("title")
		descFlagProvided := cmd.Flags().Changed("description")
		if !titleFlagProvided && !descFlagProvided {
			fmt.Println("No changes specified. Please provide at least one field to update.")
			return
		}
		// Get project's current data from db
		var (
			projectTitle       string
			projectDescription sql.NullString
			projectId          int
		)
		err := Database.QueryRow(sql_select_single_project_for_update, targetProjectId).Scan(&projectId, &projectTitle, &projectDescription)
		if err == sql.ErrNoRows {
			fmt.Printf("There is no project with the id of %d.\n", targetProjectId)
			return
		}
		if err != nil {
			fmt.Printf("Error querying project: %v.\n", err)
			return
		}
		// Set default values from existing project
		finalTitle := projectTitle
		finalDesc := ""
		if projectDescription.Valid {
			finalDesc = projectDescription.String
		}
		// Override with provided flags
		if titleFlagProvided {
			finalTitle = newProjectTitle
		}
		if descFlagProvided {
			finalDesc = newProjectDescription
		}
		// Update project
		_, err = Database.Exec(sql_update_project, finalTitle, finalDesc, time.Now(), targetProjectId)
		if err != nil {
			fmt.Printf("Error updating project: %v\n", err)
			return
		}
		// Report what was updated
		if titleFlagProvided && descFlagProvided {
			fmt.Printf("Updated both title and description for project #%d.\n", targetProjectId)
		} else if titleFlagProvided {
			fmt.Printf("Updated title for project #%d.\n", targetProjectId)
		} else if descFlagProvided {
			fmt.Printf("Updated description for project #%d.\n", targetProjectId)
		}
	},
}

func init() {
	editProject.Flags().IntVarP(&targetProjectId, "id", "i", 0, "Project id")
	editProject.Flags().StringVarP(&newProjectTitle, "title", "t", "", "Project title")
	editProject.Flags().StringVarP(&newProjectDescription, "description", "d", "", "Project description")
	editProject.MarkFlagRequired("id")
	rootCmd.AddCommand(editProject)
}
