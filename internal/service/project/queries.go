package project

import (
	"fmt"
	"os"

	"github.com/samhodg1993/toto/internal/models"
)

// ListProjects returns all projects as a slice of Project models
func (s *Service) ListProjects() ([]models.Project, error) {
	rows, err := s.db.Query("SELECT id, title, description, filepath, archived, created_at, updated_at FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		project, err := scanRowToProject(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning project: %w", err)
		}
		projects = append(projects, project)
	}

	return projects, rows.Err()
}

// GetProjectIdByFilepath returns the project ID for the current directory
func (s *Service) GetProjectIdByFilepath() (int, error) {
	var projectId int = 0

	currentDir, err := os.Getwd()
	if err != nil {
		return 0, fmt.Errorf("error getting current directory: %w", err)
	}

	row := s.db.QueryRow("SELECT id FROM projects WHERE filepath = ?", currentDir)

	err = row.Scan(&projectId)
	if err != nil {
		return 0, fmt.Errorf("no project exists for this filepath")
	}

	return projectId, nil
}
