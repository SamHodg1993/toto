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
	Description  string       `db:"description" json:"description"`
	Status       string       `db:"status" json:"status"`          // e.g., "To Do", "In Progress", "Done"
	ProjectKey   string       `db:"project_key" json:"projectKey"` // e.g., "PROJ"
	IssueType    string       `db:"issue_type" json:"issueType"`   // e.g., "Story", "Bug", "Task"
	URL          string       `db:"url" json:"url"`                // Full Jira URL
	LastSyncedAt sql.NullTime `db:"last_synced_at" json:"lastSyncedAt"`
	CreatedAt    time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time    `db:"updated_at" json:"updatedAt"`
}

// ADFNode represents a node in Atlassian Document Format
type ADFNode struct {
	Type    string    `json:"type"`
	Text    string    `json:"text,omitempty"`
	Content []ADFNode `json:"content,omitempty"`
}

// JiraBasedTicket represents a Jira ticket coming from Jira
type JiraBasedTicket struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields struct {
		Summary     string `json:"summary"`
		Description struct {
			Type    string    `json:"type"`
			Version int       `json:"version"`
			Content []ADFNode `json:"content"`
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

// extractTextFromADF recursively extracts text from ADF nodes
func extractTextFromADF(nodes []ADFNode, depth int) string {
	var text string
	for _, node := range nodes {
		switch node.Type {
		case "text":
			text += node.Text
		case "paragraph":
			text += extractTextFromADF(node.Content, depth+1) + "\n\n"
		case "bulletList", "orderedList":
			text += extractTextFromADF(node.Content, depth+1)
		case "listItem":
			// Add bullet point or number prefix
			prefix := "• "
			if depth > 0 {
				prefix = strings.Repeat("  ", depth-1) + prefix
			}
			text += prefix + strings.TrimSpace(extractTextFromADF(node.Content, depth+1)) + "\n"
		case "hardBreak":
			text += "\n"
		default:
			// For unknown types, try to extract content recursively
			if len(node.Content) > 0 {
				text += extractTextFromADF(node.Content, depth+1)
			}
		}
	}
	return text
}

func (j *JiraBasedTicket) GetDescriptionText() string {
	text := extractTextFromADF(j.Fields.Description.Content, 0)
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
