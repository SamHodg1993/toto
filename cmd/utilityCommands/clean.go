package utilityCommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CleanUtility = &cobra.Command{
	Use:   "clean",
	Short: "cls, cls-comp * ls",
	Long:  "Clears the terminal, then removes completed todos and prints the remaining",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := UtilityService.CleanAndPrintTodos()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	},
}
