package todo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	bulkCompleteSelectedIds      string = ""
	rangeBulkCompleteSelectedIds string = ""
)

var ToggleComplete = &cobra.Command{
	Use:   "toggle-complete",
	Short: "Toggle a todos status between complete and pending.",
	Long:  "Toggle an exising todos status between complete and pending.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputIdString := args[0]
		inputId, err := strconv.Atoi(inputIdString)
		if err != nil {
			fmt.Printf("Unable to parse input ID to integer type. Error: %s", err)
			return
		}

		var ids []int
		bulkCompleteLen := 0
		if len(bulkCompleteSelectedIds) > 0 {
			bulkArray :=
				strings.Split(bulkCompleteSelectedIds, ",")
			bulkCompleteLen = len(bulkArray)
		}
		rangeBulkCompleteLen := 0
		if len(rangeBulkCompleteSelectedIds) > 0 {
			rangeArray :=
				strings.Split(rangeBulkCompleteSelectedIds, ",")
			rangeBulkCompleteLen = len(rangeArray)
		}

		if bulkCompleteLen > 0 && rangeBulkCompleteLen > 0 {
			fmt.Println(`Unable to complete bulk operations with range set and integer set.
				Please pick either the range or the integer list.`)
			return
		}

		if bulkCompleteLen > 0 && rangeBulkCompleteLen < 1 {
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
		} else if bulkCompleteLen < 1 && rangeBulkCompleteLen > 0 {
			allRanges := strings.Split(rangeBulkCompleteSelectedIds, ",")

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
			ids = append(ids, inputId)
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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputIdString := args[0]
		inputId, err := strconv.Atoi(inputIdString)
		if err != nil {
			fmt.Printf("Unable to parse input ID to integer type. Error: %s", err)
			return
		}

		var ids []int
		bulkCompleteLen := 0
		if len(bulkCompleteSelectedIds) > 0 {
			bulkArray :=
				strings.Split(bulkCompleteSelectedIds, ",")
			bulkCompleteLen = len(bulkArray)
		}
		rangeBulkCompleteLen := 0
		if len(rangeBulkCompleteSelectedIds) > 0 {
			rangeArray :=
				strings.Split(rangeBulkCompleteSelectedIds, ",")
			rangeBulkCompleteLen = len(rangeArray)
		}

		if bulkCompleteLen > 0 && rangeBulkCompleteLen > 0 {
			fmt.Println(`Unable to complete bulk operations with range set and integer set.
				Please pick either the range or the integer list.`)
			return
		}

		if bulkCompleteLen > 0 && rangeBulkCompleteLen < 1 {
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
		} else if bulkCompleteLen < 1 && rangeBulkCompleteLen > 0 {
			allRanges := strings.Split(rangeBulkCompleteSelectedIds, ",")

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
			ids = append(ids, inputId)
		}

		TodoService.ToggleComplete(ids)
		if len(ids) < 2 {
			fmt.Printf("Successfully toggled the todo!\n")
		} else {
			fmt.Printf("Successfully toggled %d todos!\n", len(ids))
		}
	},
}

func init() {
	ToggleComp.Flags().StringVarP(&bulkCompleteSelectedIds, "Todo IDS", "I", "", "The target todos ID's separated with commas")
	ToggleComplete.Flags().StringVarP(&bulkCompleteSelectedIds, "Todo IDS", "I", "", "The target todos ID's separated with commas")
	ToggleComp.Flags().StringVarP(&rangeBulkCompleteSelectedIds, "Todo Range IDS", "R", "", "The target todos ID ranges separated with commas")
	ToggleComplete.Flags().StringVarP(&rangeBulkCompleteSelectedIds, "Todo Range IDS", "R", "", "The target todos ID ranges separated with commas")
}
