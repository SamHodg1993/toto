package todo

import (
	"database/sql"

	"github.com/ODGY8/toto/internal/models"
)

// ProjectServiceInterface defines methods needed from project service
type ProjectServiceInterface interface {
	GetProjectIdByFilepath() (int, error)
	HandleNoExistingProject() (int, error)
	AddNewProjectWithPrompt() error
	HandleAddNewProject(title, description string) error
}

// Service handles todo operations
type Service struct {
	db             *sql.DB
	projectService ProjectServiceInterface
}

// New creates a new todo service
func New(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// SetProjectService allows injecting the project service dependency
func (s *Service) SetProjectService(projectService ProjectServiceInterface) {
	s.projectService = projectService
}

// scanRowToTodo converts a SQL row to a Todo model (for detailed queries)
func scanRowToTodo(rows *sql.Rows) (models.Todo, error) {
	var t models.Todo
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.ProjectId, &t.CreatedAt, &t.UpdatedAt, &t.Completed, &t.CompletedAt)
	return t, err
}

// scanRowToSimpleTodo converts a SQL row to a Todo model (for simple queries: id, title, completed)
func scanRowToSimpleTodo(rows *sql.Rows) (models.Todo, error) {
	var t models.Todo
	err := rows.Scan(&t.ID, &t.Title, &t.Completed)
	return t, err
}
