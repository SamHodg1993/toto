package db

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	// Get current user
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	homeDir := currentUser.HomeDir

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Autoincrement missing is intentional to allow sqlite to reuse deleted id's
	const sql_create_todo_table = `
	CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	project_id INTEGER NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE,
	completed_at DATETIME,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
	`

	sql_create_project_table := `
	CREATE TABLE IF NOT EXISTS projects (
	id INTEGER PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	archived BOOLEAN NOT NULL DEFAULT FALSE,
	filepath VARCHAR(255) NOT NULL DEFAULT '` + filepath.Clean(homeDir) + `',
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
	`

	const sql_insert_initial_project = `
	INSERT INTO projects (title, created_at, updated_at) 
	SELECT "Global list", ?, ?
	WHERE NOT EXISTS (SELECT 1 FROM projects)
	`

	_, err = db.Exec(sql_create_todo_table)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sql_create_project_table)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sql_insert_initial_project, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	return db, nil

}
