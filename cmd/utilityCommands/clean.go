package utilityCommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	reverseList bool
)

var CleanUtility = &cobra.Command{
	Use:   "clean",
	Short: "cls, cls-comp * ls",
	Long:  "Clears the terminal, then removes completed todos and prints the remaining",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := UtilityService.CleanAndPrintTodos(reverseList)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	},
}

func init() {
	// Full date
	CleanUtility.Flags().BoolVarP(&reverseList, "reverse", "r", false, "Reverse the list of todo's. Useful when you created todos most important -> least important")
}
