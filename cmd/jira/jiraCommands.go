package jira

import (
	"fmt"

	"github.com/samhodg1993/toto-todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var JiraPull = &cobra.Command{
	Use:   "jira-pull",
	Short: "Create a new todo from a jira ticket",
	Long:  "Pull a ticket from Jira and create a new todo based on the tickets available information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todo, err := service.GetSingleJiraTicket("MBA-6")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(todo)
	},
}
