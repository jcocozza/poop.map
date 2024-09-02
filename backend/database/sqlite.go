package database

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const databasePath = "poop_locations_database.db"

// Connect to a SQLite database.
func connectToSQLite(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		slog.Error("Failed to open database")
		return nil, err
	}
	// Ping the database to ensure connectivity
	err = db.Ping()
	if err != nil {
		slog.Error("Failed to connect to database - ping failed")
		return nil, err
	}
	slog.Debug("Connected to SQLite Database")
	return db, nil
}

func createDatabaseFile(dbPath string) error {
	// Ensure the directory structure exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	// Create the database file
	_, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	slog.Debug("Database file created at: " + dbPath)
	return nil
}

// Will attempt to connect to the database that is packaged with the app.
func ConnectToDatabase() (*sql.DB, error) {
	// Check if the database file exists, if not, create it
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		if err := createDatabaseFile(databasePath); err != nil {
			return nil, err
		}
	}
	db, err1 := connectToSQLite(databasePath)
	if err1 != nil {
		return nil, err1
	}
	return db, nil
}
