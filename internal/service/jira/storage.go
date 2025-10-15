package jira

import (
	"fmt"
	"time"

	"github.com/samhodg1993/toto/internal/models"
)

// InsertJiraTicket inserts or updates a Jira ticket in the database
func (s *Service) InsertJiraTicket(ticket *models.JiraTicket) (int64, error) {
	// Check if ticket already exists
	var existingId int64
	err := s.db.QueryRow("SELECT id FROM jira_tickets WHERE jira_key = ?", ticket.JiraKey).Scan(&existingId)

	if err == nil {
		// Ticket exists, update it and return existing ID
		_, updateErr := s.db.Exec(
			`UPDATE jira_tickets
			SET title = ?, status = ?, project_key = ?, issue_type = ?, url = ?, last_synced_at = ?
			WHERE jira_key = ?`,
			ticket.Title,
			ticket.Status,
			ticket.ProjectKey,
			ticket.IssueType,
			ticket.URL,
			time.Now(),
			ticket.JiraKey,
		)
		if updateErr != nil {
			return 0, fmt.Errorf("Failed to update existing jira ticket: %v", updateErr)
		}
		return existingId, nil
	}

	// Ticket doesn't exist, insert new one
	result, err := s.db.Exec(
		`INSERT INTO jira_tickets (
		  jira_key, title, status, project_key, issue_type, url, last_synced_at
		) VALUES (?,?,?,?,?,?,?)`,
		ticket.JiraKey,
		ticket.Title,
		ticket.Status,
		ticket.ProjectKey,
		ticket.IssueType,
		ticket.URL,
		time.Now(),
	)

	if err != nil {
		return 0, fmt.Errorf("Could not insert new jira ticket, err: %v", err)
	}

	returnId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Insert Jira ticket success, but failed to get row ID, err: %v", err)
	}

	return returnId, nil
}
