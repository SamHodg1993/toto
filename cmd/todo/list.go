package todo

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/samhodg1993/toto-todo-cli/internal/utilities"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	fullDate  bool = false
	allTodos  bool = false
	clearTerm bool = false
)

var GetCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo's for current project",
	Long:  "Get a list of all the todo's for the current project (defined by the current directory).",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearTerm {
			utilities.ClearScreen()
		}

		// rows, err := service.GetTodosForFilepath()
		rows, err := TodoService.GetTodosForFilepath()
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

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

var GetCmdLong = &cobra.Command{
	Use:   "list-long",
	Short: "List todo's with more data for the current project.",
	Long:  "Get a more detailed list of all the todo's for the current project (defined by the current directory)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearTerm {
			utilities.ClearScreen()
		}

		var rows *sql.Rows
		var err error

		if allTodos {
			rows, err = TodoService.GetAllTodos_LONG()
		} else {
			rows, err = TodoService.GetTodosForFilepath_LONG()
		}

		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}
		defer rows.Close()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "Project Id", "Created At", "Updated At", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				completed   bool
				description sql.NullString
				projectId   int
				createdAt   time.Time
				updatedAt   time.Time
			)

			err := rows.Scan(&id, &title, &description, &projectId, &createdAt, &updatedAt, &completed)
			if err != nil {
				fmt.Printf("Error reading row: %v\n", err)
				return
			}

			if completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if completed {
				status = "Done"
			}

			dateFormat := "02-01-2006"
			if fullDate {
				dateFormat = time.RFC3339
			}

			table.Append([]string{
				fmt.Sprintf("%d", id),
				title,
				description.String,
				strconv.Itoa(projectId),
				createdAt.Format(dateFormat),
				updatedAt.Format(dateFormat),
				status,
			})
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List outstanding todo's",
	Long:  "Get a list of all the todo titles that are outstanding",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearTerm {
			utilities.ClearScreen()
		}

		rows, err := TodoService.GetTodosForFilepath()
		if err != nil {
			if err.Error() == "operation cancelled by user" {
				return
			}
			fmt.Printf("%v\n", err)
			return
		}

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

var LsCmdLong = &cobra.Command{
	Use:   "lsl",
	Short: "List todo's with more data",
	Long:  "Get a more detailed list of all the todo's for the current project (defined by the current directory)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearTerm {
			utilities.ClearScreen()
		}

		rows, err := TodoService.GetTodosForFilepath_LONG()
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "ProjectId", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				completed   bool
				description sql.NullString
				projectId   int
				createdAt   time.Time
				updatedAt   time.Time
				completedAt sql.NullTime
			)

			err := rows.Scan(&id, &title, &description, &projectId, &createdAt, &updatedAt, &completed, &completedAt)
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

			completedAtString := "-"
			if completedAt.Valid {
				if fullDate {
					completedAtString = completedAt.Time.Format(time.RFC3339)
				} else {
					completedAtString = completedAt.Time.Format("02-01-2006")
				}
			}

			if fullDate {
				table.Append([]string{
					fmt.Sprintf("%d", id),
					title,
					description.String,
					strconv.Itoa(projectId),
					createdAt.Format(time.RFC3339),
					updatedAt.Format(time.RFC3339),
					status,
					completedAtString,
				})
			} else {
				table.Append([]string{
					fmt.Sprintf("%d", id),
					title,
					description.String,
					strconv.Itoa(projectId),
					createdAt.Format("02-01-2006"),
					updatedAt.Format("02-01-2006"),
					status,
					completedAtString,
				})
			}
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

var LslCmdLong = &cobra.Command{
	Use:   "lsla",
	Short: "List todo's with more data regardless of project",
	Long:  "Get a more detailed list of all the todo's for all projects",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearTerm {
			utilities.ClearScreen()
		}

		rows, err := TodoService.GetAllTodos_LONG()
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "ProjectId", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for rows.Next() {
			var (
				id          int
				title       string
				completed   bool
				description sql.NullString
				projectId   int
				createdAt   time.Time
				updatedAt   time.Time
				completedAt sql.NullTime
			)

			err := rows.Scan(&id, &title, &description, &projectId, &createdAt, &updatedAt, &completed, &completedAt)
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

			completedAtString := "-"
			if completedAt.Valid {
				if fullDate {
					completedAtString = completedAt.Time.Format(time.RFC3339)
				} else {
					completedAtString = completedAt.Time.Format("02-01-2006")
				}
			}

			if fullDate {
				table.Append([]string{
					fmt.Sprintf("%d", id),
					title,
					description.String,
					strconv.Itoa(projectId),
					createdAt.Format(time.RFC3339),
					updatedAt.Format(time.RFC3339),
					status,
					completedAtString,
				})
			} else {
				table.Append([]string{
					fmt.Sprintf("%d", id),
					title,
					description.String,
					strconv.Itoa(projectId),
					createdAt.Format("02-01-2006"),
					updatedAt.Format("02-01-2006"),
					status,
					completedAtString,
				})
			}
		}

		table.Render()

		if err := rows.Err(); err != nil {
			fmt.Printf("Error iterating over rows: %v\n", err)
		}
	},
}

func init() {
	// Full date
	LsCmdLong.Flags().BoolVarP(&fullDate, "Full-Date", "D", false, "Return the dates as full timestamps")
	GetCmdLong.Flags().BoolVarP(&fullDate, "Full-Date", "D", false, "Return the dates as full timestamps")
	LslCmdLong.Flags().BoolVarP(&fullDate, "Full-Date", "D", false, "Return the dates as full timestamps")
	// All todos
	GetCmdLong.Flags().BoolVarP(&allTodos, "All-Todos", "A", false, "Return all todo's regardless of project")
	// Clear before print
	GetCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	LsCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	LsCmdLong.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	GetCmdLong.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	LslCmdLong.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
}
