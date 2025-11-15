package projects

import (
	"fmt"
	"os"

	"github.com/ODGY8/toto/internal/utilities"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var clearTerm bool = false

var ProjectLsCmd = &cobra.Command{
	Use:   "proj-ls",
	Short: "List project's",
	Long:  "Get a list of all the projects titles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := ProjectService.ListProjects()
		if err != nil {
			fmt.Printf("There was an error getting the project's from the database: %v\n", err)
			return
		}

		if clearTerm {
			utilities.ClearScreen()
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Project Title", "Description", "Filepath", "Archived"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, project := range projects {
			title := project.Title
			// If project has been archived, apply strikethrough to the title
			if project.Archived {
				title = strikethrough(title)
			}

			status := "Active"
			if project.Archived {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", project.ID),
				title,
				project.Description,
				project.Filepath,
				status,
			})
		}

		table.Render()
	},
}

var ProjectListCmd = &cobra.Command{
	Use:   "project-list",
	Short: "List project's",
	Long:  "Get a list of all the projects titles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := ProjectService.ListProjects()
		if err != nil {
			fmt.Printf("There was an error getting the project's from the database: %v\n", err)
			return
		}

		if clearTerm {
			utilities.ClearScreen()
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Project Title", "Description", "Filepath", "Archived"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, project := range projects {
			title := project.Title
			// If project has been archived, apply strikethrough to the title
			if project.Archived {
				title = strikethrough(title)
			}

			status := "Active"
			if project.Archived {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", project.ID),
				title,
				project.Description,
				project.Filepath,
				status,
			})
		}

		table.Render()
	},
}

func init() {
	ProjectLsCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the project list")
	ProjectListCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the project list")
}
