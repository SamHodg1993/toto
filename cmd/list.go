package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var sql_get_todos string = "SELECT id, title, created_at, completed FROM todos"

var getCmd = &cobra.Command{
	Use:   "list",
	Short: "List outstanding todo's",
	Long:  "Get a list of all the todo titles that are outstanding",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := database.Query(sql_get_todos)
		if err != nil {
			fmt.Printf("There was an error getting the todo's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Created At", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var id int
			var title string
			var createdAt string
			var completed bool

			err := rows.Scan(&id, &title, &createdAt, &completed)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}

			// If todo is completed, apply strikethrough to the title
			if completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if completed {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", id),
				title,
				createdAt,
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
	rootCmd.AddCommand(getCmd)
}
