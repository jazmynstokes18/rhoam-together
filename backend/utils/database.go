package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB creates a connection to the PostgreSQL database
// To be fully implemented in Phase 3
func ConnectDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("✓ Database connection successful")
	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *sql.DB) error {
	return db.Close()
}
