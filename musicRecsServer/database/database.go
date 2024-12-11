package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() (*sql.DB, error) {
	// Open SQLite database
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	// Create blacklist table
	blacklistTable := `
	CREATE TABLE IF NOT EXISTS blacklist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		machine_id TEXT NOT NULL,
		artist_name TEXT NOT NULL,
		UNIQUE (machine_id, artist_name)
	);`

	_, err = db.Exec(blacklistTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddToBlacklist(db *sql.DB, machineID, artistName string) error {
	query := `
	INSERT INTO blacklist (machine_id, artist_name)
	VALUES (?, ?)
	ON CONFLICT (machine_id, artist_name) DO NOTHING;
	`
	_, err := db.Exec(query, machineID, artistName)
	if err != nil {
		return fmt.Errorf("failed to add artist to blacklist: %v", err)
	}
	return nil
}

func RemoveFromBlacklist(db *sql.DB, machineID, artistName string) error {
	query := `
	DELETE FROM blacklist
	WHERE machine_id = ? AND artist_name = ?;
	`
	_, err := db.Exec(query, machineID, artistName)
	if err != nil {
		return fmt.Errorf("failed to remove artist from blacklist: %v", err)
	}
	return nil
}

func FetchBlacklist(db *sql.DB, machineID string) ([]string, error) {
	query := `
	SELECT artist_name FROM blacklist
	WHERE machine_id = ?;
	`

	rows, err := db.Query(query, machineID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch blacklist: %v", err)
	}
	defer rows.Close()

	var artists []string
	for rows.Next() {
		var artist string
		if err := rows.Scan(&artist); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		artists = append(artists, artist)
	}
	return artists, nil
}
