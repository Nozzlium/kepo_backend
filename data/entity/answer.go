package entity

import (
	"database/sql"
	"time"
)

type Answer struct {
	ID         uint
	QuestionID uint
	UserID     uint
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}
