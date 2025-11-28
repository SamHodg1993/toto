package todo

import (
	"fmt"
	"strconv"

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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetIdString := args[0]
		targetId, err := strconv.Atoi(targetIdString)
		if err != nil {
			fmt.Printf("Unable to parse input ID to integer type. Error: %s", err)
			return
		}

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
	EditTodo.Flags().StringVarP(&newTitle, "title", "t", "", "Todo title")
	EditTodo.Flags().StringVarP(&newDescription, "description", "d", "", "Todo description")

	EditTodo.MarkFlagRequired("id")
}
