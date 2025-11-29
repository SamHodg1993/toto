package jira

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/odgy8/toto/internal/models"
)

func (s *Service) HandleListJiraTickets() error {
	// Get the jira tickets data
	rows, err := s.db.Query(`
          SELECT id, jira_key, title, description, status,
                 project_key, issue_type, url, last_synced_at,
                 created_at, updated_at
          FROM jira_tickets
          ORDER BY created_at DESC
      `)
	if err != nil {
		return fmt.Errorf("failed to query: %v", err)
	}
	defer rows.Close()

	// Scan into slice
	var tickets []models.JiraTicket
	for rows.Next() {
		var ticket models.JiraTicket
		err := rows.Scan(
			&ticket.ID,
			&ticket.JiraKey,
			&ticket.Title,
			&ticket.Description,
			&ticket.Status,
			&ticket.ProjectKey,
			&ticket.IssueType,
			&ticket.URL,
			&ticket.LastSyncedAt,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("error scanning ticket: %w", err)
		}
		tickets = append(tickets, ticket)
	}

	// Build and render table here
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Jira Key", "Title", "Status", "Type"})

	for _, ticket := range tickets {
		table.Append([]string{
			strconv.Itoa(ticket.ID),
			ticket.JiraKey,
			ticket.Title,
			ticket.Status,
			ticket.IssueType,
		})
	}

	table.Render()
	return nil
}
