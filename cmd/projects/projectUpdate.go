package projects

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/samhodg1993/toto-todo-cli/cmd"

	"github.com/spf13/cobra"
)

var sql_update_project string = "UPDATE projects SET title = ?, description = ?, filepath = ?, updated_at = ? WHERE id = ?"
var sql_select_single_project_for_update string = "SELECT id, title, description, filepath FROM projects WHERE id = ?"
var (
	targetProjectId       int
	newProjectDescription string
	newProjectTitle       string
	newFilepath           string
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
		filepathFlagProvided := cmd.Flags().Changed("filepath")
		if !titleFlagProvided && !descFlagProvided && !filepathFlagProvided {
			fmt.Println("No changes specified. Please provide at least one field to update.")
			return
		}

		// Get project's current data from db
		var (
			projectTitle       string
			projectDescription sql.NullString
			projectId          int
			filepath           string
		)
		err := cmd.Database.QueryRow(sql_select_single_project_for_update, targetProjectId).Scan(&projectId, &projectTitle, &projectDescription, &filepath)
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
		finalFilepath := filepath
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

		if filepathFlagProvided {
			finalFilepath = newFilepath
		}

		// Update project
		_, err = cmd.Database.Exec(sql_update_project, finalTitle, finalDesc, finalFilepath, time.Now(), targetProjectId)
		if err != nil {
			fmt.Printf("Error updating project: %v\n", err)
			return
		}

		// Report what was updated
		if titleFlagProvided && descFlagProvided && filepathFlagProvided {
			fmt.Printf("Updated title, description and filepath for project #%d.\n", targetProjectId)
		} else if titleFlagProvided && descFlagProvided {
			fmt.Printf("Updated title and description for project #%d.\n", targetProjectId)
		} else if titleFlagProvided && filepathFlagProvided {
			fmt.Printf("Updated title and filepath for project #%d.\n", targetProjectId)
		} else if descFlagProvided && filepathFlagProvided {
			fmt.Printf("Updated description and filepath for project #%d.\n", targetProjectId)
		} else if titleFlagProvided {
			fmt.Printf("Updated title for project #%d.\n", targetProjectId)
		} else if descFlagProvided {
			fmt.Printf("Updated description for project #%d.\n", targetProjectId)
		} else if filepathFlagProvided {
			fmt.Printf("Updated filepath for project #%d.\n", targetProjectId)
		}
	},
}

func init() {
	editProject.Flags().IntVarP(&targetProjectId, "id", "i", 0, "Project id")
	editProject.Flags().StringVarP(&newProjectTitle, "title", "t", "", "Project title")
	editProject.Flags().StringVarP(&newProjectDescription, "description", "d", "", "Project description")
	editProject.Flags().StringVarP(&newFilepath, "filepath", "f", "", "Projects filepath")

	editProject.MarkFlagRequired("id")

	cmd.RootCmd.AddCommand(editProject)
}
