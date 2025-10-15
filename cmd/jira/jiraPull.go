package jira

import (
	"fmt"
	"time"

	"github.com/samhodg1993/toto-todo-cli/cmd/projects"
	"github.com/samhodg1993/toto-todo-cli/cmd/todo"
	"github.com/samhodg1993/toto-todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var jiraTicketId string

var JiraPull = &cobra.Command{
	Use:   "jira-pull",
	Short: "Create a new todo from a jira ticket",
	Long:  "Pull a ticket from Jira and create a new todo based on the tickets available information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if jiraTicketId == "" {
			fmt.Println("You must provide this function with a jira ticket id using the -i flag.")
			return
		}

		ticket, err := JiraService.GetSingleJiraTicket(jiraTicketId)
		if err != nil {
			fmt.Println(err)
			return
		}

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
			fmt.Printf("Failed to insert jira ticket. Err: %v", err)
			return
		}

		projectId, err := projects.ProjectService.GetProjectIdByFilepath()
		if err != nil {
			fmt.Printf("Error occured when collecting filepath project when pulling jira ticket. Exited with error: %v", err)
			return
		}

		if err = todo.TodoService.AddTodo(
			jiraTicket.Title,
			ticket.GetDescriptionText(),
			projectId,
			time.Now(),
			time.Now(),
			jiraInsertId,
		); err != nil {
			fmt.Printf("Failed to store new Todo, jira table row created. Error: %v", err)
			return
		}

		fmt.Printf("Successfully pulled Jira ticket %s and created todo!\n", ticket.Key)
	},
}

func init() {
	JiraPull.PersistentFlags().StringVarP(&jiraTicketId, "jiraTicketId", "i", "", "Jira id for the ticket")
}
