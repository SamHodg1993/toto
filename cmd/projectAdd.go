package cmd

import (
	"fmt"
	"os"

	"github.com/samhodg1993/todo-cli/models"

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
		var project models.NewProject

		// Set the project details
		project.Title = projectTitle
		project.Description = projectDescription

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
		// Set the filepath
		project.Filepath = currentDir

		// Add the new project
		AddNewProject(project)

		fmt.Printf("New project added: %s\n", projectTitle)
	},
}

var projAddCmd = &cobra.Command{
	Use:   "proj-add",
	Short: "Add a new project.",
	Long:  "Add a new project to the list of stored projects.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var project models.NewProject

		// Set the project details
		project.Title = projectTitle
		project.Description = projectDescription

		// Get the current working directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		// Set the project filepath
		if projectFilepath == "" {
			project.Filepath = currentDir
			fmt.Printf("Filepath not provided defaulting to working directory: %s\n", currentDir)
		} else {
			project.Filepath = projectFilepath
		}

		// Add the project
		AddNewProject(project)
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
