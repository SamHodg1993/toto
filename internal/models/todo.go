package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Todo struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	ProjectId   int       `db:"project_id" json:"projectId"`
	Completed   bool      `db:"completed" json:"completed"`
	CompletedAt sql.NullTime `db:"completed_at" json:"completedAt"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

// IsValid checks if a todo has valid data
func (t *Todo) IsValid() bool {
	return t.Title != "" && t.ProjectId > 0
}

// IsCompleted returns whether the todo is completed
func (t *Todo) IsCompleted() bool {
	return t.Completed
}

// TimeAgo returns a human-readable time since creation
func (t *Todo) TimeAgo() string {
	if t.CreatedAt.IsZero() {
		return "unknown"
	}

	duration := time.Since(t.CreatedAt)
	if duration.Hours() < 24 {
		return fmt.Sprintf("%.0f hours ago", duration.Hours())
	} else {
		return fmt.Sprintf("%.0f days ago", duration.Hours()/24)
	}
}
