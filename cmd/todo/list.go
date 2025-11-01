package todo

import (
	"fmt"
	"os"
	"slices"

	"github.com/samhodg1993/toto/internal/models"
	todoHelper "github.com/samhodg1993/toto/internal/service/todo"
	"github.com/samhodg1993/toto/internal/utilities"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	fullDate    bool = false
	allTodos    bool = false
	clearTerm   bool = false
	reverseList bool = false
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

		if reverseList {
			slices.Reverse(todos)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			table.Append(todoHelper.FormatTodoTableRow(todo, strikethrough))
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

		if reverseList {
			slices.Reverse(todos)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "Project Id", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			table.Append(todoHelper.FormatTodoTableRowLong(todo, fullDate, strikethrough))
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

		if reverseList {
			slices.Reverse(todos)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Status"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			table.Append(todoHelper.FormatTodoTableRow(todo, strikethrough))
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

		if reverseList {
			slices.Reverse(todos)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "ProjectId", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			table.Append(todoHelper.FormatTodoTableRowLong(todo, fullDate, strikethrough))
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

		if reverseList {
			slices.Reverse(todos)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Todo", "Description", "ProjectId", "Created At", "Updated At", "Status", "Completed At"})
		table.SetBorder(true)
		table.SetRowLine(true)

		strikethrough := color.New(color.CrossedOut).SprintFunc()

		for _, todo := range todos {
			table.Append(todoHelper.FormatTodoTableRowLong(todo, fullDate, strikethrough))
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
	// Reversr order (userful for when there are laods of todo's and they were created important to least important)
	GetCmd.Flags().BoolVarP(&reverseList, "Reverse the todo order", "r", false, "Reverse the list of todo's. Useful when you created todos most important -> least important")
	LsCmd.Flags().BoolVarP(&reverseList, "Reverse the todo order", "r", false, "Reverse the list of todo's. Useful when you created todos most important -> least important")
	LsCmdLong.Flags().BoolVarP(&reverseList, "Reverse the todo order", "r", false, "Reverse the list of todo's. Useful when you created todos most important -> least important")
	GetCmdLong.Flags().BoolVarP(&reverseList, "Reverse the todo order", "r", false, "Reverse the list of todo's. Useful when you created todos most important -> least important")
	LslCmdLong.Flags().BoolVarP(&reverseList, "Reverse the todo order", "r", false, "Reverse the list of todo's. Useful when you created todos most important -> least important")
}
