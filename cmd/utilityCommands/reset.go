package utilityCommands

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/samhodg1993/toto/internal/service/database"

	"github.com/spf13/cobra"
)

var (
	confirmFlag bool
	dbService   *database.Service
)

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database",
	Long:  "Reset the database, remove all existing todos and projects and set the ids back to 1",
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

		// Use the dbService to reset the database
		err := dbService.ResetDatabase()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Database cleared successfully!")
	},
}

// InitDBService initializes the DB service with a database connection
func InitDBService(db *sql.DB) {
	dbService = database.New(db)
}

func init() {
	ResetCmd.Flags().BoolVarP(&confirmFlag, "confirm", "c", false, "Skip confirmation prompt")
}
