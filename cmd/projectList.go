package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var sql_get_projects string = "SELECT id, title, filepath, archived FROM projects"

var projectLsCmd = &cobra.Command{
	Use:   "prls",
	Short: "List project's",
	Long:  "Get a list of all the projects titles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := Database.Query(sql_get_projects)
		if err != nil {
			fmt.Printf("There was an error getting the project's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Project Title", "Filepath", "Archived"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id       int
				title    string
				archived bool
				filepath string
			)

			err := rows.Scan(&id, &title, &filepath, &archived)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}

			// If project has been archived, apply strikethrough to the title
			if archived {
				title = strikethrough(title)
			}

			status := "Pending"
			if archived {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", id),
				title,
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

var projectListCmd = &cobra.Command{
	Use:   "project-list",
	Short: "List project's",
	Long:  "Get a list of all the projects titles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := Database.Query(sql_get_projects)
		if err != nil {
			fmt.Printf("There was an error getting the project's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Project Title", "Filepath", "Archived"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id       int
				title    string
				archived bool
				filepath string
			)

			err := rows.Scan(&id, &title, &filepath, &archived)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}

			// If project has been archived, apply strikethrough to the title
			if archived {
				title = strikethrough(title)
			}

			status := "Pending"
			if archived {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", id),
				title,
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
	rootCmd.AddCommand(projectLsCmd)
	rootCmd.AddCommand(projectListCmd)
}
