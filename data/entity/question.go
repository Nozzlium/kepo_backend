package entity

import (
	"database/sql"
	"time"
)

type Question struct {
	ID          uint
	UserID      uint
	CategoryID  uint
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Description string
}
