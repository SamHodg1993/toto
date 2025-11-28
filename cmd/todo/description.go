package todo

import (
	"fmt"
	"strconv"

	"github.com/odgy8/toto/internal/utilities"
	"github.com/spf13/cobra"
)

var (
	descriptionSelectedTodo int  = 0
	clearScreen             bool = false
)

var GetTodoDescription = &cobra.Command{
	Use:   "description",
	Short: "Get the description for a single todo.",
	Long:  "Get the description for a single todo.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if clearScreen {
			utilities.ClearScreen()
		}

		descriptionId := args[0]
		stringId, err := strconv.Atoi(descriptionId)
		if err != nil {
			fmt.Printf("Failed to parse input ID into integer. Error: %s", err)
		}

		desc, err := TodoService.GetTodoDetails(stringId)
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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if clearScreen {
			utilities.ClearScreen()
		}

		descriptionId := args[0]
		stringId, err := strconv.Atoi(descriptionId)
		if err != nil {
			fmt.Printf("Failed to parse input ID into integer. Error: %s", err)
		}

		desc, err := TodoService.GetTodoDetails(stringId)
		if err != nil {
			fmt.Printf("%v.\n", err)
			return
		}

		fmt.Printf("%v\n", desc)
	},
}

func init() {
	GetTodoDescription.Flags().BoolVarP(&clearScreen, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
	GetTodoDesc.Flags().BoolVarP(&clearScreen, "Clear terminal first", "C", false, "Clear the terminal before listing the todos")
}
