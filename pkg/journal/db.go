package journal

import (
	"database/sql"
)

// Initialize db if not already initialized
func dbInit(db *sql.DB) error {
	const schema = `PRAGMA foreign_keys = ON;
					CREATE TABLE IF NOT EXISTS users (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						name TEXT NOT NULL
					);
					CREATE TABLE IF NOT EXISTS goals (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						user INTEGER,
						title TEXT NOT NULL,
						created INTEGER NOT NULL,
						status INTEGER NOT NULL,
						FOREIGN KEY (user) REFERENCES users (id)
					);
					CREATE TABLE IF NOT EXISTS responses (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						goal INTEGER,
						created INTEGER NOT NULL,
						value INTEGER NOT NULL,
						FOREIGN KEY (goal) REFERENCES goals (id)
					);
					`
	_, err := db.Exec(schema)
	return err
}

// Adds a new user to the db
func addUser(db *sql.DB, userName string) (*User, error) {
	const insertQuery = `INSERT INTO users (name) VALUES (?)`
	result, err := db.Exec(insertQuery, userName)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &User{id, userName}, err
}
