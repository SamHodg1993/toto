package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sql_get_todos string = "SELECT title FROM todos"

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

		fmt.Println("List of todo's:")
		for rows.Next() {
			var title string
			err := rows.Scan(&title)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}
			fmt.Printf("- %s\n", title)
		}

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
