package entity

import "time"

type Notification struct {
	ID         int
	UserID     int
	QuestionID int
	NotifType  string
	Headline   string
	Preview    string
	IsRead     bool
	CreatedAt  time.Time
}
