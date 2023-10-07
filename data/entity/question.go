package entity

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID          uint
	UserID      uint
	CategoryID  uint
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Description string
}
