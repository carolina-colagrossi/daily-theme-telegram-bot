package journal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewJournal(file string) (*Journal, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	return &Journal{db}, dbInit(db)
}

// Adds a new user to the journal
func (journal *Journal) AddNewUser(user string) (*User, error) {
	return addUser(journal.db, user)
}
