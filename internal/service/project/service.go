package project

import (
	"database/sql"

	"github.com/samhodg1993/toto/internal/models"
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
func scanRowToProject(rows *sql.Rows) (models.Project, error) {
	var p models.Project
	err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Filepath, &p.Archived, &p.CreatedAt, &p.UpdatedAt, &p.JiraURL)
	return p, err
}
