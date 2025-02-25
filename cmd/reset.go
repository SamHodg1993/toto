package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var sqlDeleteAllTodos string = "DELETE FROM todos"
var sqlResetAutoIncrement string = "DELETE FROM sqlite_sequence WHERE name='todos'"
var confirmFlag bool

var resetTodos = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database",
	Long:  "Reset the database, remove all existing todos and set the id back to 1",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if --confirm flag was provided
		if !confirmFlag {
			// If not confirmed via flag, ask for confirmation
			fmt.Print("You are about to remove all data from the database. Please confirm that you want to continue (y/N): ")
			var userInput string
			fmt.Scanln(&userInput)

			if strings.ToLower(userInput) != "y" && strings.ToLower(userInput) != "yes" {
				fmt.Println("Operation cancelled. Aborting!")
				return
			}
		}

		// Use a transaction to ensure both operations complete
		tx, err := database.Begin()
		if err != nil {
			fmt.Printf("Error starting transaction: %v\n", err)
			return
		}

		// Delete all todos
		_, err = tx.Exec(sqlDeleteAllTodos)
		if err != nil {
			tx.Rollback()
			fmt.Printf("Error clearing the database: %v\n", err)
			return
		}

		// Reset the auto-increment counter
		_, err = tx.Exec(sqlResetAutoIncrement)
		if err != nil {
			tx.Rollback()
			fmt.Printf("Error resetting ID sequence: %v\n", err)
			return
		}

		// Commit the transaction
		err = tx.Commit()
		if err != nil {
			fmt.Printf("Error committing transaction: %v\n", err)
			return
		}

		fmt.Println("Database cleared successfully!")
	},
}

func init() {
	resetTodos.Flags().BoolVarP(&confirmFlag, "confirm", "c", false, "Skip confirmation prompt")

	rootCmd.AddCommand(resetTodos)
}
