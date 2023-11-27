package entity

import "time"

type Notification struct {
	ID         uint
	UserID     uint
	QuestionID uint
	NotifType  string
	Headline   string
	Preview    string
	IsRead     bool
	CreatedAt  time.Time
}
