package todo

import (
	"fmt"

	"github.com/samhodg1993/toto/internal/utilities"
	"github.com/spf13/cobra"
)

var (
	descriptionSelectedTodo int
	clearScreen             bool = false
)

var GetTodoDescription = &cobra.Command{
	Use:   "description",
	Short: "Get the description for a single todo.",
	Long:  "Get the description for a single todo.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearScreen {
			utilities.ClearScreen()
		}

		desc, err := TodoService.GetTodoDetails(descriptionSelectedTodo)
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		fmt.Printf("%v\n", desc)
	},
}

var GetTodoDesc = &cobra.Command{
	Use:   "desc",
	Short: "Get the description for a single todo.",
	Long:  "Get the description for a single todo.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if clearScreen {
			utilities.ClearScreen()
		}

		desc, err := TodoService.GetTodoDetails(descriptionSelectedTodo)
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		fmt.Printf("%v\n", desc)
	},
}

func init() {
	GetTodoDescription.Flags().IntVarP(&descriptionSelectedTodo, "Todo ID", "i", 0, "The target todo's ID")
	GetTodoDesc.Flags().IntVarP(&descriptionSelectedTodo, "Todo ID", "i", 0, "The target todo's ID")
	GetTodoDescription.Flags().BoolVarP(&clearScreen, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	GetTodoDesc.Flags().BoolVarP(&clearScreen, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
}
