package service

import (
	"database/sql"
	"time"
)

// User represents user's account.
type User struct {
	ID          int64
	Username    sql.NullString
	FirstName   string
	LastName    string
	BirthDay    sql.NullTime
	Gender      sql.NullString
	Email       string
	PhoneNumber sql.NullString
	UpdatedAt   sql.NullTime
	CreatedAt   time.Time
}
