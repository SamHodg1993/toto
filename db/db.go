package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "todo.db"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	const sql_create_table = `
	CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = db.Exec(sql_create_table)
	if err != nil {
		return nil, err
	}

	return db, nil
}
