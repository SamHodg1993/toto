package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"path/filepath"
)

var devMode = os.Getenv("IS_DEV_MODE")
var dbPath string

func init() {
	devMode := os.Getenv("IS_DEV_MODE")
	if devMode != "" {
		dbPath = "./todos.db"
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Errorf("Could not find home directory: %v", err)
		}
		dbPath = filepath.Join(homeDir, ".todos.db")
	}
}

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
	completed BOOLEAN NOT NULL DEFAULT FALSE,
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
