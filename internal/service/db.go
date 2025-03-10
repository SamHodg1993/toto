package service

import (
	"database/sql"
	"fmt"
)

// DBService handles database-wide operations
type DBService struct {
	db *sql.DB
}

// NewDBService creates a new database service
func NewDBService(db *sql.DB) *DBService {
	return &DBService{db: db}
}

// ResetDatabase clears all data from the database and resets auto-increment values
func (s *DBService) ResetDatabase() error {
	// Use a transaction to ensure all operations complete successfully
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	// Delete all todos
	_, err = tx.Exec("DELETE FROM todos")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error clearing the todo database: %w", err)
	}

	// Delete all projects
	_, err = tx.Exec("DELETE FROM projects")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error clearing the projects database: %w", err)
	}

	// Reset the auto-increment counter for todos
	_, err = tx.Exec("DELETE FROM sqlite_sequence WHERE name='todos'")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error resetting todo ID sequence: %w", err)
	}

	// Reset the auto-increment counter for projects
	_, err = tx.Exec("DELETE FROM sqlite_sequence WHERE name='projects'")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error resetting project ID sequence: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
