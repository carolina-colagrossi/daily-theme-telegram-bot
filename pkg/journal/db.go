package journal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Initialize db if not already initialized
func dbInit(db *sql.DB) {
	const schema = ""
}
