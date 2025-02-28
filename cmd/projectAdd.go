package cmd

import (
	"github.com/spf13/cobra"
)

var (
	projectTitle       string
	projectDescription string
	projectFilepath    string
)

var projectAddCmd = &cobra.Command{
	Use:   "project-add",
	Short: "Add a new project.",
	Long:  "Add a new project to the list of stored projects.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		HandleAddNewProject(projectTitle, projectDescription)
	},
}

var projAddCmd = &cobra.Command{
	Use:   "proj-add",
	Short: "Add a new project.",
	Long:  "Add a new project to the list of stored projects.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		HandleAddNewProject(projectTitle, projectDescription)
	},
}

func init() {
	projectAddCmd.PersistentFlags().StringVarP(&projectTitle, "title", "t", "", "Title of the project")
	projectAddCmd.PersistentFlags().StringVarP(&projectDescription, "description", "d", "", "Description of the project")
	projectAddCmd.PersistentFlags().StringVarP(&projectFilepath, "filepath", "f", "", "Set the file path for the project")

	projAddCmd.PersistentFlags().StringVarP(&projectTitle, "title", "t", "", "Title of the project")
	projAddCmd.PersistentFlags().StringVarP(&projectDescription, "description", "d", "", "Description of the project")
	projAddCmd.PersistentFlags().StringVarP(&projectFilepath, "filepath", "f", "", "Set the file path for the project")

	projectAddCmd.MarkFlagRequired("title")
	projAddCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(projectAddCmd)
	rootCmd.AddCommand(projAddCmd)
}
