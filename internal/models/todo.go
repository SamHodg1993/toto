package models

import (
	"time"
)

type Todo struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	ProjectId   int       `db:"project_id" json:"projectId"`
	Completed   bool      `db:"completed" json:"completed"`
	CompletedAt time.Time `db:"completed_at" json:"completedAt"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}
