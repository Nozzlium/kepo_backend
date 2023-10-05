package entity

import "time"

type Answer struct {
	ID          uint
	QuestionID  uint
	UserID      uint
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	AnswerLikes []AnswerLike
}
