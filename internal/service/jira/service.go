package jira

import (
	"database/sql"
)

// Service handles all Jira-related operations
type Service struct {
	db             *sql.DB
	todoService    TodoServiceInterface
	projectService ProjectServiceInterface
	jiraService    JiraServiceInterface
}

// New creates a new Jira service
func New(db *sql.DB) *Service {
	return &Service{db: db}
}
