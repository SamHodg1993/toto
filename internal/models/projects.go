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
}

type NewProject struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Archived    bool   `db:"archived" json:"archived"`
	Filepath    string `db:"filepath" json:"filepath"`
}
