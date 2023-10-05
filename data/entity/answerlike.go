package entity

import "time"

type AnswerLike struct {
	AnswerID  uint
	UserID    uint
	CreatedAt time.Time
}
