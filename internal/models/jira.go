package models

import (
	"database/sql"
	"time"
)

// JiraTicket represents a Jira ticket stored in our database
type JiraTicket struct {
	ID           int          `db:"id" json:"id"`
	JiraKey      string       `db:"jira_key" json:"jiraKey"` // e.g., "PROJ-123"
	Title        string       `db:"title" json:"title"`
	Status       string       `db:"status" json:"status"`          // e.g., "To Do", "In Progress", "Done"
	ProjectKey   string       `db:"project_key" json:"projectKey"` // e.g., "PROJ"
	IssueType    string       `db:"issue_type" json:"issueType"`   // e.g., "Story", "Bug", "Task"
	URL          string       `db:"url" json:"url"`                // Full Jira URL
	LastSyncedAt sql.NullTime `db:"last_synced_at" json:"lastSyncedAt"`
	CreatedAt    time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time    `db:"updated_at" json:"updatedAt"`
}

// IsValid checks if a Jira ticket has valid data
func (j *JiraTicket) IsValid() bool {
	return j.JiraKey != "" && j.Title != "" && j.URL != ""
}

// IsStatusDone checks if the Jira ticket is completed
func (j *JiraTicket) IsStatusDone() bool {
	return j.Status == "Done" || j.Status == "Closed" || j.Status == "Resolved"
}

// JiraConfig represents Jira configuration settings
type JiraConfig struct {
	BaseURL      string `json:"baseUrl"`      // e.g., "https://company.atlassian.net"
	ClientID     string `json:"clientId"`     // OAuth Client ID
	ClientSecret string `json:"clientSecret"` // OAuth Client Secret (from env)
	AccessToken  string `json:"accessToken"`  // User's access token
	ProjectKey   string `json:"projectKey"`   // Default project key
}

// IsConfigured checks if Jira is properly configured
func (j *JiraConfig) IsConfigured() bool {
	return j.BaseURL != "" && j.AccessToken != "" && j.ProjectKey != ""
}
