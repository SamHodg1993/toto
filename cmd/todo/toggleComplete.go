package todo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	completeSelectedId      int    = 0
	bulkCompleteSelectedIds string = ""
)

var ToggleComplete = &cobra.Command{
	Use:   "toggle-complete",
	Short: "Toggle a todos status between complete and pending.",
	Long:  "Toggle an exising todos status between complete and pending.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		if len(bulkCompleteSelectedIds) < 1 {
			ids = append(ids, completeSelectedId)
		} else {
			allIds := strings.Split(bulkCompleteSelectedIds, ",")

			for _, id := range allIds {
				id = strings.TrimSpace(id)
				asInt, err := strconv.Atoi(id)

				if err != nil {
					fmt.Printf("Unable to convert ID to integer for ID: %s. Skipping...\n", id)
					continue
				}

				ids = append(ids, asInt)
			}
		}

		TodoService.ToggleComplete(ids)
		if len(ids) < 2 {
			fmt.Printf("Successfully toggled the todo!\n")
		} else {
			fmt.Printf("Successfully toggled %d todos!\n", len(ids))
		}
	},
}

var ToggleComp = &cobra.Command{
	Use:   "comp",
	Short: "Toggle a todos status between complete and pending.",
	Long:  "Toggle an exising todos status between complete and pending.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		if len(bulkCompleteSelectedIds) < 1 {
			ids = append(ids, completeSelectedId)
		} else {
			allIds := strings.Split(bulkCompleteSelectedIds, ",")

			for _, id := range allIds {
				id = strings.TrimSpace(id)
				asInt, err := strconv.Atoi(id)

				if err != nil {
					fmt.Printf("Unable to convert ID to integer for ID: %s. Skipping...\n", id)
					continue
				}

				ids = append(ids, asInt)
			}
		}

		TodoService.ToggleComplete(ids)
	},
}

func init() {
	ToggleComp.Flags().IntVarP(&completeSelectedId, "Todo ID", "i", 0, "The target todos ID")
	ToggleComplete.Flags().IntVarP(&completeSelectedId, "Todo ID", "i", 0, "The target todos ID")
	ToggleComp.Flags().StringVarP(&bulkCompleteSelectedIds, "Todo IDS", "I", "", "The target todos ID's separated with commas")
	ToggleComplete.Flags().StringVarP(&bulkCompleteSelectedIds, "Todo IDS", "I", "", "The target todos ID's separated with commas")
}
