package projects

import (
	"fmt"
	"os"

	"github.com/samhodg1993/toto-todo-cli/internal/utilities"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var sql_get_projects string = "SELECT id, title, filepath, archived FROM projects"

var clearTerm bool = false

var ProjectLsCmd = &cobra.Command{
	Use:   "proj-ls",
	Short: "List project's",
	Long:  "Get a list of all the projects titles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := ProjectService.ListProjects()
		if err != nil {
			fmt.Printf("There was an error getting the project's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		if clearTerm {
			utilities.ClearScreen()
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Project Title", "Description", "Filepath", "Archived"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				description string
				filepath    string
				archived    bool
			)

			err := rows.Scan(&id, &title, &description, &filepath, &archived)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}

			// If project has been archived, apply strikethrough to the title
			if archived {
				title = strikethrough(title)
			}

			status := "Active"
			if archived {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", id),
				title,
				description,
				filepath,
				status,
			})
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

var ProjectListCmd = &cobra.Command{
	Use:   "project-list",
	Short: "List project's",
	Long:  "Get a list of all the projects titles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := ProjectService.ListProjects()
		if err != nil {
			fmt.Printf("There was an error getting the project's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		if clearTerm {
			utilities.ClearScreen()
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Project Title", "Descrition", "Filepath", "Archived"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				description string
				filepath    string
				archived    bool
			)

			err := rows.Scan(&id, &title, &description, &filepath, &archived)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}

			// If project has been archived, apply strikethrough to the title
			if archived {
				title = strikethrough(title)
			}

			status := "Active"
			if archived {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", id),
				title,
				description,
				filepath,
				status,
			})
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

func init() {
	ProjectLsCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the project list")
	ProjectListCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the project list")
}
