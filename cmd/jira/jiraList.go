package jira

import (
	"fmt"

	"github.com/spf13/cobra"
)

var JiraList = &cobra.Command{
	Use:   "jira-list",
	Short: "List all Jira tickets",
	Long:  "List all tickets from Jira that are stored in the database.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := JiraService.HandleListJiraTickets()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var JL = &cobra.Command{
	Use:   "jl",
	Short: "List all Jira tickets",
	Long:  "List all tickets from Jira that are stored in the database.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := JiraService.HandleListJiraTickets()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
