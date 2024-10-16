package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Book struct {
	ID     string `json:"id"` // Use string for easier storage in SQLite
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Server struct {
	DB *sql.DB // SQLite database connection
}

// Initialize the AppState and create the books table if it doesn't exist
func NewServer(dbPath string) (*Server, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT,
		author TEXT
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return &Server{DB: db}, nil
}
