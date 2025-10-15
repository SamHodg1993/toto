package jira

import (
	"fmt"

	"github.com/spf13/cobra"
)

var jiraTicketId string

var JiraPull = &cobra.Command{
	Use:   "jira-pull",
	Short: "Create a new todo from a jira ticket",
	Long:  "Pull a ticket from Jira and create a new todo based on the tickets available information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := JiraService.HandlePullTicket(jiraTicketId); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	JiraPull.PersistentFlags().StringVarP(&jiraTicketId, "jiraTicketId", "i", "", "Jira id for the ticket")
}
