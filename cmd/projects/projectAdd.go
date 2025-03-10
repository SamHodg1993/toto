package projects

import (
	"github.com/spf13/cobra"
)

var (
	projectTitle       string
	projectDescription string
	projectFilepath    string
)

var ProjectAddCmd = &cobra.Command{
	Use:   "project-add",
	Short: "Add a new project.",
	Long:  "Add a new project to the list of stored projects.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ProjectService.HandleAddNewProject(projectTitle, projectDescription)
	},
}

var ProjAddCmd = &cobra.Command{
	Use:   "proj-add",
	Short: "Add a new project.",
	Long:  "Add a new project to the list of stored projects.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ProjectService.HandleAddNewProject(projectTitle, projectDescription)
	},
}

func init() {
	ProjectAddCmd.PersistentFlags().StringVarP(&projectTitle, "title", "t", "", "Title of the project")
	ProjectAddCmd.PersistentFlags().StringVarP(&projectDescription, "description", "d", "", "Description of the project")
	ProjectAddCmd.PersistentFlags().StringVarP(&projectFilepath, "filepath", "f", "", "Set the file path for the project")

	ProjAddCmd.PersistentFlags().StringVarP(&projectTitle, "title", "t", "", "Title of the project")
	ProjAddCmd.PersistentFlags().StringVarP(&projectDescription, "description", "d", "", "Description of the project")
	ProjAddCmd.PersistentFlags().StringVarP(&projectFilepath, "filepath", "f", "", "Set the file path for the project")

	ProjectAddCmd.MarkFlagRequired("title")
	ProjAddCmd.MarkFlagRequired("title")
}
