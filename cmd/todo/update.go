package todo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	targetId       int
	newDescription string
	newTitle       string
)

var EditTodo = &cobra.Command{
	Use:   "edit",
	Short: "Update a todo",
	Long:  "Update a todo's title or description",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if ID is provided
		if targetId == 0 {
			fmt.Println("-i is a required flag.")
			return
		}

		// Check if any update flags were provided
		titleFlagProvided := cmd.Flags().Changed("title")
		descFlagProvided := cmd.Flags().Changed("description")

		if !titleFlagProvided && !descFlagProvided {
			fmt.Println("No changes specified. Please provide at least one field to update.")
			return
		}

		message, err := TodoService.UpdateTodo(targetId, newTitle, newDescription, titleFlagProvided, descFlagProvided)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(message)
	},
}

func init() {
	EditTodo.Flags().IntVarP(&targetId, "id", "i", 0, "Todo id")
	EditTodo.Flags().StringVarP(&newTitle, "title", "t", "", "Todo title")
	EditTodo.Flags().StringVarP(&newDescription, "description", "d", "", "Todo description")

	EditTodo.MarkFlagRequired("id")
}
