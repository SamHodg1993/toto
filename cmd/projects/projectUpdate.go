package projects

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	targetProjectId       int
	newProjectTitle       string
	newProjectDescription string
	newFilepath           string
)

var EditProject = &cobra.Command{
	Use:   "edit",
	Short: "Update a project",
	Long:  "Update a project's title, description, or filepath",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if ID is provided
		if targetProjectId == 0 {
			fmt.Println("-i or --id is a required flag.")
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

		// Call the service to update the project
		message, err := ProjectService.UpdateProject(
			targetProjectId,
			newProjectTitle,
			newProjectDescription,
			newFilepath,
			titleFlagProvided,
			descFlagProvided,
			filepathFlagProvided,
		)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(message)
	},
}

func init() {
	EditProject.Flags().IntVarP(&targetProjectId, "id", "i", 0, "ID of the project to update")
	EditProject.Flags().StringVarP(&newProjectTitle, "title", "t", "", "New title for the project")
	EditProject.Flags().StringVarP(&newProjectDescription, "description", "d", "", "New description for the project")
	EditProject.Flags().StringVarP(&newFilepath, "filepath", "f", "", "New filepath for the project")

	EditProject.MarkFlagRequired("id")
}
