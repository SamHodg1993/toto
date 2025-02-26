package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var sql_insert_project string = "INSERT INTO projects (title, description, archived, filepath, created_at, updated_at) VALUES (?,?,?,?,?,?)"

var (
	projectTitle       string
	projectDescription string
	projectCreatedAt   string
	projectUpdatedAt   string
	projectArchived    int
	projectFilepath    string
)

var projectAddCmd = &cobra.Command{
	Use:   "project-add",
	Short: "Add a new project.",
	Long:  "Add a new project to the list of stored projects.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Create default variables
		createdAt := time.Now()
		updatedAt := time.Now()
		filepath := "~/"

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		if projectFilepath == "" {
			filepath = currentDir
			fmt.Printf("Filepath not provided defaulting to working directory: %s", currentDir)
		} else {
			filepath = projectFilepath
		}

		// If created at flag has value, update the created at value
		if todoCreatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoCreatedAt)
			if err == nil {
				createdAt = parsed
			}
		}

		// If updated at flag has value, update the created at value
		if todoUpdatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoUpdatedAt)
			if err == nil {
				updatedAt = parsed
			}
		}

		// Update the todo
		_, err = database.Exec(sql_insert_project, projectTitle, projectDescription, projectArchived, filepath, createdAt, updatedAt)
		if err != nil {
			fmt.Printf("There was an error adding the todo: %v\n", err)
			return
		}
		fmt.Printf("New todo added: %s\n", todoTitle)
	},
}

func init() {
	projectAddCmd.PersistentFlags().StringVarP(&projectTitle, "title", "t", "", "Title of the todo")
	projectAddCmd.PersistentFlags().StringVarP(&projectDescription, "description", "d", "", "Description of the todo")
	projectAddCmd.PersistentFlags().StringVarP(&projectCreatedAt, "created-at", "c", "", "Todo creation time")
	projectAddCmd.PersistentFlags().StringVarP(&projectUpdatedAt, "updated-at", "u", "", "Todo last updated time")
	projectAddCmd.PersistentFlags().IntVarP(&projectArchived, "archived", "a", 0, "Relevent Project Id")
	projectAddCmd.PersistentFlags().StringVarP(&projectFilepath, "filepath", "f", "", "Set the file path for the project")

	addCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(projectAddCmd)
}
