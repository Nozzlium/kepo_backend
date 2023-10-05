package entity

import "time"

type Question struct {
	ID          uint
	UserID      uint
	CategoryID  uint
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
}
