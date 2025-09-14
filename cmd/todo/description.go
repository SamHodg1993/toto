package todo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	descriptionSelectedTodo int
)

var GetTodoDescription = &cobra.Command{
	Use:   "description",
	Short: "Get the description for a single todo.",
	Long:  "Get the description for a single todo.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// rows, err := service.GetTodosForFilepath()
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
		// rows, err := service.GetTodosForFilepath()
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
}
