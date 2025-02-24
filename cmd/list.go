package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var sql_get_todos string = "SELECT id, title, completed FROM todos"
var sql_get_todos_LONG string = "SELECT id, title, description, created_at, updated_at, completed FROM todos"

var getCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo's",
	Long:  "Get a list of all the todo titles and completed status.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := database.Query(sql_get_todos)
		if err != nil {
			fmt.Printf("There was an error getting the todo's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id        int
				title     string
				completed bool
			)

			err := rows.Scan(&id, &title, &completed)
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
				status,
			})
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

var getCmdLong = &cobra.Command{
	Use:   "list-long",
	Short: "List todo's with more data",
	Long:  "Get a list of all the todo titles, descriptions, completed status, created at and upated at timestamps",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := database.Query(sql_get_todos_LONG)
		if err != nil {
			fmt.Printf("There was an error getting the todo's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "Created At", "Updated At", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				completed   bool
				description sql.NullString
				createdAt   string
				updatedAt   string
			)

			err := rows.Scan(&id, &title, &description, &createdAt, &updatedAt, &completed)
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
				description.String,
				createdAt,
				updatedAt,
				status,
			})
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List outstanding todo's",
	Long:  "Get a list of all the todo titles that are outstanding",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := database.Query(sql_get_todos)
		if err != nil {
			fmt.Printf("There was an error getting the todo's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id        int
				title     string
				completed bool
			)

			err := rows.Scan(&id, &title, &completed)
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
				status,
			})
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

var lsCmdLong = &cobra.Command{
	Use:   "lsla",
	Short: "List todo's with more data",
	Long:  "Get a list of all the todo titles, descriptions, completed status, created at and upated at timestamps",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := database.Query(sql_get_todos_LONG)
		if err != nil {
			fmt.Printf("There was an error getting the todo's from the database: %v\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "Created At", "Updated At", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				completed   bool
				description sql.NullString
				createdAt   string
				updatedAt   string
			)

			err := rows.Scan(&id, &title, &description, &createdAt, &updatedAt, &completed)
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
				description.String,
				createdAt,
				updatedAt,
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
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(getCmdLong)
	rootCmd.AddCommand(lsCmdLong)
}
