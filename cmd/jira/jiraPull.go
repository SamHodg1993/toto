package jira

import (
	"fmt"
	"time"

	"github.com/samhodg1993/toto-todo-cli/cmd/projects"
	"github.com/samhodg1993/toto-todo-cli/cmd/todo"
	"github.com/samhodg1993/toto-todo-cli/internal/service"
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

		ticket, err := service.GetSingleJiraTicket(jiraTicketId)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Collate the data used by both inserts
		createdAt := time.Now()
		updatedAt := time.Now()
		projectId, err := projects.ProjectService.GetProjectIdByFilepath()
		if err != nil {
			fmt.Printf("Error occured when collecting filepath project when pulling jira ticket. Exited with error: %v", err)
		}
		title := ticket.Fields.Summary
		desc := ticket.GetDescriptionText()
		// Make a service to add a jira ticket to the jira ticket table here

		// Once added the jira ticket table row, add the todo and reference the jira ticket table row
		err = todo.TodoService.AddTodo(title, desc)
	},
}

func init() {
	JiraPull.PersistentFlags().StringVarP(&jiraTicketId, "jiraTicketId", "i", "", "Jira id for the ticket")
}
