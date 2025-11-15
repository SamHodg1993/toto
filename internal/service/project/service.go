package project

import (
	"database/sql"

	"github.com/odgy8/toto/internal/models"
)

var sql_insert_project string = `
	INSERT INTO projects (
		title,
		description,
		archived,
		filepath,
		created_at,
		updated_at,
	  jira_url
	) VALUES (?,?,?,?,?,?,?)`

// Service handles project operations
type Service struct {
	db *sql.DB
}

// New creates a new project service
func New(db *sql.DB) *Service {
	return &Service{db: db}
}

// scanRowToProject converts a SQL row to a Project model
func scanRowToProject(rows *sql.Rows) (models.Project,
	error) {
	var p models.Project
	var description, jiraURL sql.NullString
	err := rows.Scan(&p.ID, &p.Title, &description,
		&p.Filepath, &p.Archived, &p.CreatedAt, &p.UpdatedAt,
		&jiraURL)

	// Convert NullString to regular string
	if description.Valid {
		p.Description = description.String
	}
	if jiraURL.Valid {
		p.JiraURL = jiraURL.String
	}

	return p, err
}
