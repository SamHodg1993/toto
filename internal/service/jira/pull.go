package jira

import (
	"fmt"
	"time"

	"github.com/odgy8/toto/internal/models"
	"github.com/odgy8/toto/internal/service/claude"
)

// TodoServiceInterface defines methods needed from todo service
type TodoServiceInterface interface {
	AddTodo(title, description string, projectId int, createdAt, updatedAt time.Time, jiraTicketId int64) error
}

// ProjectServiceInterface defines methods needed from project service
type ProjectServiceInterface interface {
	GetProjectIdByFilepath() (int, error)
	GetProjectJiraURL() (string, error)
}

// SetDependencies allows injecting todo and project services
func (s *Service) SetDependencies(todoService TodoServiceInterface, projectService ProjectServiceInterface) {
	s.todoService = todoService
	s.projectService = projectService
}

// HandlePullTicket pulls a Jira ticket and creates a single todo
func (s *Service) HandlePullTicket(ticketId string) error {
	if ticketId == "" {
		return fmt.Errorf("You must provide a jira ticket id")
	}

	// Fetch ticket from Jira
	ticket, err := s.GetSingleJiraTicket(ticketId)
	if err != nil {
		return err
	}

	// Save to database
	jiraTicket := &models.JiraTicket{
		JiraKey:    ticket.Key,
		Title:      ticket.Fields.Summary,
		Status:     ticket.Fields.Status.Name,
		ProjectKey: ticket.Fields.Project.Key,
		IssueType:  ticket.Fields.IssueType.Name,
		URL:        ticket.Self,
	}

	jiraInsertId, err := s.InsertJiraTicket(jiraTicket)
	if err != nil {
		return fmt.Errorf("Failed to insert jira ticket: %v", err)
	}

	// Get project ID
	projectId, err := s.projectService.GetProjectIdByFilepath()
	if err != nil {
		return fmt.Errorf("Error getting project by filepath: %v", err)
	}

	// Create todo
	if err = s.todoService.AddTodo(
		jiraTicket.Title,
		ticket.GetDescriptionText(),
		projectId,
		time.Now(),
		time.Now(),
		jiraInsertId,
	); err != nil {
		return fmt.Errorf("Failed to store new Todo: %v", err)
	}

	fmt.Printf("Successfully pulled Jira ticket %s and created todo!\n", ticket.Key)
	return nil
}

// HandlePullTicketWithClaude pulls a Jira ticket and uses Claude AI to break it into subtasks
func (s *Service) HandlePullTicketWithClaude(ticketId string) error {
	if ticketId == "" {
		return fmt.Errorf("You must provide a jira ticket id")
	}

	// Fetch Jira ticket
	fmt.Printf("Fetching Jira ticket %s...\n", ticketId)
	ticket, err := s.GetSingleJiraTicket(ticketId)
	if err != nil {
		return fmt.Errorf("Error fetching Jira ticket: %v", err)
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

	jiraInsertId, err := s.InsertJiraTicket(jiraTicket)
	if err != nil {
		return fmt.Errorf("Failed to insert jira ticket: %v", err)
	}

	// Get project ID
	projectId, err := s.projectService.GetProjectIdByFilepath()
	if err != nil {
		return fmt.Errorf("Error getting project by filepath: %v", err)
	}

	// Call Claude to break down ticket into subtasks
	fmt.Println("Asking Claude to break down the ticket into subtasks...")
	descriptionText := ticket.GetDescriptionText()
	subtasks, err := claude.BreakdownJiraTicketWithClaude(
		ticket.Key,
		ticket.Fields.Summary,
		descriptionText,
	)
	if err != nil {
		return fmt.Errorf("Error calling Claude AI: %v", err)
	}

	if len(subtasks) == 0 {
		fmt.Println("Claude didn't generate any subtasks. Creating main ticket as todo instead...")
		// Fallback: create the main ticket as a todo
		if err = s.todoService.AddTodo(
			jiraTicket.Title,
			ticket.GetDescriptionText(),
			projectId,
			time.Now(),
			time.Now(),
			jiraInsertId,
		); err != nil {
			return fmt.Errorf("Failed to store todo: %v", err)
		}
		fmt.Printf("Created main todo for Jira ticket %s\n", ticket.Key)
		return nil
	}

	fmt.Printf("Creating %d subtasks...\n", len(subtasks))
	successCount := 0
	for i, subtask := range subtasks {
		if err = s.todoService.AddTodo(
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

	return nil
}
