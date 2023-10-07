package entity

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID         uint
	QuestionID uint
	UserID     uint
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
