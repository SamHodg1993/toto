package todo

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/spf13/cobra"
)

var (
	deleteSelectedId       int    = 0
	bulkDeleteSelectedIds  string = ""
	rangeDeleteSelectedIds string = ""
)

var DeleteTodo = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id using the -i flag",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		bulkDeleteLen := 0
		if len(bulkDeleteSelectedIds) > 0 {
			bulkArray :=
				strings.Split(bulkDeleteSelectedIds, ",")
			bulkDeleteLen = len(bulkArray)
		}

		rangeDeleteLen := 0
		if len(rangeDeleteSelectedIds) > 0 {
			rangeArray :=
				strings.Split(rangeDeleteSelectedIds, ",")
			rangeDeleteLen = len(rangeArray)
		}

		if bulkDeleteLen > 0 && rangeDeleteLen > 0 {
			fmt.Println(`Unable to delete with both range set and integer set.
              Please pick either the range or the integer list.`)
			return
		}

		if bulkDeleteLen > 0 && rangeDeleteLen < 1 {
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
		} else if bulkDeleteLen < 1 && rangeDeleteLen > 0 {
			allRanges := strings.Split(rangeDeleteSelectedIds, ",")

			for _, rng := range allRanges {
				perimiters := strings.Split(rng, "-")

				start, err := strconv.Atoi(perimiters[0])
				if err != nil {
					fmt.Printf("Invalid start range %s", err)
				}

				end, err := strconv.Atoi(perimiters[len(perimiters)-1])
				if err != nil {
					fmt.Printf("Invalid end range %s", err)
				}

				if start == 0 || end == 0 {
					fmt.Println("Unable to find start and end of range. Aborting operation.")
					continue
				}

				for start < end+1 {
					ids = append(ids, start)
					start++
				}
			}
		} else {
			ids = append(ids, deleteSelectedId)
		}

		TodoService.DeleteTodo(ids)
		if len(ids) < 2 {
			fmt.Printf("Successfully deleted the todo!\n")
		} else {
			fmt.Printf("Successfully deleted %d todos!\n", len(ids))
		}
	},
}

var DelTodo = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo",
	Long:  "Delete a single todo from the database by referencing the todo id using the -i flag",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		bulkDeleteLen := 0
		if len(bulkDeleteSelectedIds) > 0 {
			bulkArray :=
				strings.Split(bulkDeleteSelectedIds, ",")
			bulkDeleteLen = len(bulkArray)
		}
		rangeDeleteLen := 0
		if len(rangeDeleteSelectedIds) > 0 {
			rangeArray :=
				strings.Split(rangeDeleteSelectedIds,
					",")
			rangeDeleteLen = len(rangeArray)
		}

		if bulkDeleteLen > 0 && rangeDeleteLen > 0 {
			fmt.Println(`Unable to delete with both range
   set and integer set.
              Please pick either the range or the
  integer list.`)
			return
		}

		if bulkDeleteLen > 0 && rangeDeleteLen < 1 {
			allIds :=
				strings.Split(bulkDeleteSelectedIds, ",")

			for _, id := range allIds {
				id = strings.TrimSpace(id)
				asInt, err := strconv.Atoi(id)

				if err != nil {
					fmt.Printf("Unable to convert ID to integer for ID: %s. Skipping...\n", id)
					continue
				}

				ids = append(ids, asInt)
			}
		} else if bulkDeleteLen < 1 && rangeDeleteLen > 0 {
			allRanges := strings.Split(rangeDeleteSelectedIds, ",")

			for _, rng := range allRanges {
				perimiters := strings.Split(rng, "-")

				start, err := strconv.Atoi(perimiters[0])
				if err != nil {
					fmt.Printf("Invalid start range %s", err)
				}

				end, err := strconv.Atoi(perimiters[len(perimiters)-1])
				if err != nil {
					fmt.Printf("Invalid end range %s", err)
				}

				if start == 0 || end == 0 {
					fmt.Println("Unable to find start and end of range. Aborting operation.")
					continue
				}

				for start < end+1 {
					ids = append(ids, start)
					start++
				}
			}
		} else {
			ids = append(ids, deleteSelectedId)
		}

		TodoService.DeleteTodo(ids)
		if len(ids) < 2 {
			fmt.Printf("Successfully deleted the todo!\n")
		} else {
			fmt.Printf("Successfully deleted %d todos!\n", len(ids))
		}
	},
}

func init() {
	DeleteTodo.Flags().IntVarP(&deleteSelectedId, "Todo ID", "i", 0, "The target todos ID")
	DelTodo.Flags().IntVarP(&deleteSelectedId, "Todo ID", "i", 0, "The target todos ID")
	DeleteTodo.Flags().StringVarP(&bulkDeleteSelectedIds, "Todo IDs", "I", "", "The comma separated list of todo IDs")
	DelTodo.Flags().StringVarP(&bulkDeleteSelectedIds, "Todo IDs", "I", "", "The comma separated list of todo IDs")
	DeleteTodo.Flags().StringVarP(&rangeDeleteSelectedIds, "Todo Range IDs", "R", "", "The comma separated list of todo IDs")
	DelTodo.Flags().StringVarP(&rangeDeleteSelectedIds, "Todo Range IDs", "R", "", "The comma separated list of todo IDs")
}
