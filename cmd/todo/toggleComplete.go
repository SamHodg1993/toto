package todo

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeSelectedId int = 0

var ToggleComplete = &cobra.Command{
	Use:   "toggle-complete",
	Short: "Toggle a todo's status between complete and pending.",
	Long:  "Toggle an exising todo's status between complete and pending.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		idString := strconv.Itoa(completeSelectedId)

		newStatus, err := TodoService.ToggleComplete(idString)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		statusText := "incomplete"
		if newStatus {
			statusText = "complete"
		}

		fmt.Printf("Todo with id: %s status updated successfully to %s\n", idString, statusText)
	},
}

var ToggleComp = &cobra.Command{
	Use:   "comp",
	Short: "Toggle a todo's status between complete and pending.",
	Long:  "Toggle an exising todo's status between complete and pending.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		idString := strconv.Itoa(completeSelectedId)

		newStatus, err := TodoService.ToggleComplete(idString)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		statusText := "incomplete"
		if newStatus {
			statusText = "complete"
		}

		fmt.Printf("Todo with id: %s status updated successfully to %s\n", idString, statusText)
	},
}

func init() {
	ToggleComp.Flags().IntVarP(&completeSelectedId, "Todo ID", "i", 0, "The target todo's ID")
	ToggleComplete.Flags().IntVarP(&completeSelectedId, "Todo ID", "i", 0, "The target todo's ID")
}
