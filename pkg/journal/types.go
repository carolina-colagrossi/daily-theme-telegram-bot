package journal

import (
	"time"
	"database/sql"
)

// Journal encapsulates a database
// and provides a typed API to retrieve
// data.
type Journal struct {
	db *sql.DB
}

// User represents a user
// account
type User struct {
	Id int
	Name string
}

type GoalStatus int

const (
	GoalStatusActive GoalStatus = iota
	GoalStatusPaused
	GoalStatusDeleted
)

// Goal is a goal tracked by a user
type Goal struct {
	Id int
	user int
	Title string
	Created time.Time
	Status GoalStatus
}

type ResponseValue int

const (
	ResponseValueYes ResponseValue = iota
	ResponseValueNo
)

// Response encodes a users response to
// a daily goal
type Response struct {
	Id int
	goal int
	Created time.Time
	Value ResponseValue
}
