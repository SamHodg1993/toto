package models

import (
	"time"
)

type Project struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Archived    bool      `db:"archived" json:"archived"`
	Filepath    string    `db:"filepath" json:"filepath"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
	JiraURL     string    `db:"jira_url" json:"jiraUrl"`
}

type NewProject struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Archived    bool   `db:"archived" json:"archived"`
	Filepath    string `db:"filepath" json:"filepath"`
}

// IsValid checks if a project has valid data
func (p *Project) IsValid() bool {
	return p.Title != "" && p.Filepath != ""
}

// IsArchived returns whether the project is archived
func (p *Project) IsArchived() bool {
	return p.Archived
}
