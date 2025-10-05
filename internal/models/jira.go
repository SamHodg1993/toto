package models

import (
	"database/sql"
	"strings"
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

// JiraBasedTicket represents a Jira ticket coming from Jira
type JiraBasedTicket struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields struct {
		Summary     string `json:"summary"`
		Description struct {
			Type    string `json:"type"`
			Version int    `json:"version"`
			Content []struct {
				Type    string `json:"type"`
				Content []struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"content"`
			} `json:"content"`
		} `json:"description"`
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
		IssueType struct {
			Name string `json:"name"`
		} `json:"issuetype"`
		Project struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"project"`
	} `json:"fields"`
}

func (j *JiraBasedTicket) GetDescriptionText() string {
	var text string
	for _, paragraph := range j.Fields.Description.Content {
		for _, content := range paragraph.Content {
			text += content.Text + " "
		}
	}
	return strings.TrimSpace(text)
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

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}
