package jira

import (
	"fmt"

	"github.com/spf13/cobra"
)

var jiraTicketClaudeId string

var JiraPullClaude = &cobra.Command{
	Use:   "jira-pull-claude",
	Short: "Pull Jira ticket and use Claude AI to break it into subtasks",
	Long:  "Pull a ticket from Jira, use Claude AI to break it down into actionable subtasks, and create todos for each subtask linked to the Jira ticket.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := JiraService.HandlePullTicketWithClaude(jiraTicketClaudeId); err != nil {
			fmt.Println(err)
			return
		}
	},
}

var JPC = &cobra.Command{
	Use:   "jpc",
	Short: "Pull Jira ticket and use Claude AI to break it into subtasks",
	Long:  "Pull a ticket from Jira, use Claude AI to break it down into actionable subtasks, and create todos for each subtask linked to the Jira ticket.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := JiraService.HandlePullTicketWithClaude(jiraTicketClaudeId); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	JiraPullClaude.PersistentFlags().StringVarP(&jiraTicketClaudeId, "jiraTicketId", "i", "", "Jira id for the ticket")
	JPC.PersistentFlags().StringVarP(&jiraTicketClaudeId, "jiraTicketId", "i", "", "Jira id for the ticket")
}
