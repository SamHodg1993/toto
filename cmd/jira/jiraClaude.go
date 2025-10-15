package jira

import (
	"fmt"
	"time"

	"github.com/samhodg1993/toto-todo-cli/cmd/projects"
	"github.com/samhodg1993/toto-todo-cli/cmd/todo"
	"github.com/samhodg1993/toto-todo-cli/internal/models"
	"github.com/samhodg1993/toto-todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var jiraTicketClaudeId string

var JiraPullClaude = &cobra.Command{
	Use:   "jira-pull-claude",
	Short: "Pull Jira ticket and use Claude AI to break it into subtasks",
	Long:  "Pull a ticket from Jira, use Claude AI to break it down into actionable subtasks, and create todos for each subtask linked to the Jira ticket.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if jiraTicketClaudeId == "" {
			fmt.Println("You must provide this function with a jira ticket id using the -i flag.")
			return
		}

		// Fetch Jira ticket
		fmt.Printf("Fetching Jira ticket %s...\n", jiraTicketClaudeId)
		ticket, err := JiraService.GetSingleJiraTicket(jiraTicketClaudeId)
		if err != nil {
			fmt.Printf("Error fetching Jira ticket: %v\n", err)
			return
		}

		// Save Jira ticket to database
		jiraTicket := &models.JiraTicket{
			JiraKey:    ticket.Key,
			Title:      ticket.Fields.Summary,
			Status:     ticket.Fields.Status.Name,
			ProjectKey: ticket.Fields.Project.Key,
			IssueType:  ticket.Fields.IssueType.Name,
			URL:        ticket.Self,
		}

		jiraInsertId, err := JiraService.InsertJiraTicket(jiraTicket)
		if err != nil {
			fmt.Printf("Failed to insert jira ticket. Err: %v\n", err)
			return
		}

		// Get project ID
		projectId, err := projects.ProjectService.GetProjectIdByFilepath()
		if err != nil {
			fmt.Printf("Error occurred when collecting filepath project. Error: %v\n", err)
			return
		}

		// Call Claude to break down ticket into subtasks
		fmt.Println("Asking Claude to break down the ticket into subtasks...")
		descriptionText := ticket.GetDescriptionText()
		subtasks, err := service.BreakdownJiraTicketWithClaude(
			ticket.Key,
			ticket.Fields.Summary,
			descriptionText,
		)
		if err != nil {
			fmt.Printf("Error calling Claude AI: %v\n", err)
			return
		}

		if len(subtasks) == 0 {
			fmt.Println("Claude didn't generate any subtasks. Creating main ticket as todo instead...")
			// Fallback: create the main ticket as a todo
			if err = todo.TodoService.AddTodo(
				jiraTicket.Title,
				ticket.GetDescriptionText(),
				projectId,
				time.Now(),
				time.Now(),
				jiraInsertId,
			); err != nil {
				fmt.Printf("Failed to store todo. Error: %v\n", err)
				return
			}
			fmt.Printf("Created main todo for Jira ticket %s\n", ticket.Key)
			return
		}

		fmt.Printf("Creating %d subtasks...\n", len(subtasks))
		successCount := 0
		for i, subtask := range subtasks {
			if err = todo.TodoService.AddTodo(
				subtask.Title,
				subtask.Description,
				projectId,
				time.Now(),
				time.Now(),
				jiraInsertId,
			); err != nil {
				fmt.Printf("Warning: Failed to create subtask %d (%s): %v\n", i+1, subtask.Title, err)
			} else {
				successCount++
			}
		}

		fmt.Printf("\nSuccessfully pulled Jira ticket %s and created %d/%d subtasks!\n", ticket.Key, successCount, len(subtasks))
		fmt.Println("\nSubtasks created:")
		for i, subtask := range subtasks {
			fmt.Printf("  %d. %s\n", i+1, subtask.Title)
		}
	},
}

func init() {
	JiraPullClaude.PersistentFlags().StringVarP(&jiraTicketClaudeId, "jiraTicketId", "i", "", "Jira id for the ticket")
}
