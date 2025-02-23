package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sql_insert_todo string = "INSERT INTO todos (title) VALUES (?)"

var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a new todo",
	Long:  "Add a new todo to the list of stored todos",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := args[0]
		_, err := database.Exec(sql_insert_todo, todo)
		if err != nil {
			fmt.Printf("There was an error adding the todo: %v\n", err)
			return
		}
		fmt.Printf("New todo added: %s\n", todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
