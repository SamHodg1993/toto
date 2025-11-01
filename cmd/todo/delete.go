package todo

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/spf13/cobra"
)

var (
	deleteSelectedId      int    = 0
	bulkDeleteSelectedIds string = ""
)

var DeleteTodo = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id using the -i flag",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		if len(bulkDeleteSelectedIds) < 1 {
			ids = append(ids, deleteSelectedId)
		} else {
			allIds := strings.Split(bulkDeleteSelectedIds, ",")

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

		TodoService.DeleteTodo(ids)
	},
}

var DelTodo = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id using the -i flag",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		if len(bulkDeleteSelectedIds) < 1 {
			ids = append(ids, deleteSelectedId)
		} else {
			allIds := strings.Split(bulkDeleteSelectedIds, ",")

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

		TodoService.DeleteTodo(ids)
	},
}

func init() {
	DeleteTodo.Flags().IntVarP(&deleteSelectedId, "Todo ID", "i", 0, "The target todos ID")
	DelTodo.Flags().IntVarP(&deleteSelectedId, "Todo ID", "i", 0, "The target todos ID")
	DeleteTodo.Flags().StringVarP(&bulkDeleteSelectedIds, "Todo IDs", "I", "", "The comma separated list of todo IDs")
	DelTodo.Flags().StringVarP(&bulkDeleteSelectedIds, "Todo IDs", "I", "", "The comma separated list of todo IDs")
}
