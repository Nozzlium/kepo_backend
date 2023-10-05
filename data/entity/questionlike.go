package entity

import "time"

type QuestionLike struct {
	UserID     uint
	QuestionID uint
	CreatedAt  time.Time
}
