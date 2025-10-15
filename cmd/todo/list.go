package todo

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/samhodg1993/toto/internal/models"
	"github.com/samhodg1993/toto/internal/utilities"

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

		var todos []models.Todo
		var err error

		if allTodos {
			todos, err = TodoService.GetAllTodos()
		} else {
			todos, err = TodoService.GetTodosForFilepath()
		}

		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			title := todo.Title
			// If todo is completed, apply strikethrough to the title
			if todo.Completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if todo.Completed {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", todo.ID),
				title,
				status,
			})
		}

		table.Render()
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

		var todos []models.Todo
		var err error

		if allTodos {
			todos, err = TodoService.GetAllTodos_LONG()
		} else {
			todos, err = TodoService.GetTodosForFilepath_LONG()
		}

		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "Project Id", "Created At", "Updated At", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			title := todo.Title
			if todo.Completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if todo.Completed {
				status = "Done"
			}

			dateFormat := "02-01-2006"
			if fullDate {
				dateFormat = time.RFC3339
			}

			table.Append([]string{
				fmt.Sprintf("%d", todo.ID),
				title,
				todo.Description,
				strconv.Itoa(todo.ProjectId),
				todo.CreatedAt.Format(dateFormat),
				todo.UpdatedAt.Format(dateFormat),
				status,
			})
		}

		table.Render()
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

		var todos []models.Todo
		var err error

		if allTodos {
			todos, err = TodoService.GetAllTodos()
		} else {
			todos, err = TodoService.GetTodosForFilepath()
		}

		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			title := todo.Title
			// If todo is completed, apply strikethrough to the title
			if todo.Completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if todo.Completed {
				status = "Done"
			}

			table.Append([]string{
				fmt.Sprintf("%d", todo.ID),
				title,
				status,
			})
		}

		table.Render()
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

		var todos []models.Todo
		var err error

		if allTodos {
			todos, err = TodoService.GetAllTodos_LONG()
		} else {
			todos, err = TodoService.GetTodosForFilepath_LONG()
		}

		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "ProjectId", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			title := todo.Title
			// If todo is completed, apply strikethrough to the title
			if todo.Completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if todo.Completed {
				status = "Done"
			}

			completedAtString := "-"
			if todo.CompletedAt.Valid {
				if fullDate {
					completedAtString = todo.CompletedAt.Time.Format(time.RFC3339)
				} else {
					completedAtString = todo.CompletedAt.Time.Format("02-01-2006")
				}
			}

			if fullDate {
				table.Append([]string{
					fmt.Sprintf("%d", todo.ID),
					title,
					todo.Description,
					strconv.Itoa(todo.ProjectId),
					todo.CreatedAt.Format(time.RFC3339),
					todo.UpdatedAt.Format(time.RFC3339),
					status,
					completedAtString,
				})
			} else {
				table.Append([]string{
					fmt.Sprintf("%d", todo.ID),
					title,
					todo.Description,
					strconv.Itoa(todo.ProjectId),
					todo.CreatedAt.Format("02-01-2006"),
					todo.UpdatedAt.Format("02-01-2006"),
					status,
					completedAtString,
				})
			}
		}

		table.Render()
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

		todos, err := TodoService.GetAllTodos_LONG()
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "ProjectId", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			title := todo.Title
			// If todo is completed, apply strikethrough to the title
			if todo.Completed {
				title = strikethrough(title)
			}

			status := "Pending"
			if todo.Completed {
				status = "Done"
			}

			completedAtString := "-"
			if todo.CompletedAt.Valid {
				if fullDate {
					completedAtString = todo.CompletedAt.Time.Format(time.RFC3339)
				} else {
					completedAtString = todo.CompletedAt.Time.Format("02-01-2006")
				}
			}

			if fullDate {
				table.Append([]string{
					fmt.Sprintf("%d", todo.ID),
					title,
					todo.Description,
					strconv.Itoa(todo.ProjectId),
					todo.CreatedAt.Format(time.RFC3339),
					todo.UpdatedAt.Format(time.RFC3339),
					status,
					completedAtString,
				})
			} else {
				table.Append([]string{
					fmt.Sprintf("%d", todo.ID),
					title,
					todo.Description,
					strconv.Itoa(todo.ProjectId),
					todo.CreatedAt.Format("02-01-2006"),
					todo.UpdatedAt.Format("02-01-2006"),
					status,
					completedAtString,
				})
			}
		}

		table.Render()
	},
}

func init() {
	// Full date
	LsCmdLong.Flags().BoolVarP(&fullDate, "Full-Date", "D", false, "Return the dates as full timestamps")
	GetCmdLong.Flags().BoolVarP(&fullDate, "Full-Date", "D", false, "Return the dates as full timestamps")
	LslCmdLong.Flags().BoolVarP(&fullDate, "Full-Date", "D", false, "Return the dates as full timestamps")
	// All todos
	GetCmd.Flags().BoolVarP(&allTodos, "All-Todos", "A", false, "Return all todo's regardless of project")
	GetCmdLong.Flags().BoolVarP(&allTodos, "All-Todos", "A", false, "Return all todo's regardless of project")
	LsCmd.Flags().BoolVarP(&allTodos, "All-Todos", "A", false, "Return all todo's regardless of project")
	LsCmdLong.Flags().BoolVarP(&allTodos, "All-Todos", "A", false, "Return all todo's regardless of project")
	// Clear before print
	GetCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	LsCmd.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	LsCmdLong.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	GetCmdLong.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	LslCmdLong.Flags().BoolVarP(&clearTerm, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
}
