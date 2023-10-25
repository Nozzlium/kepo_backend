package result

import "time"

type QuestionResult struct {
	ID           uint      `gorm:"column:id"`
	Content      string    `gorm:"column:content"`
	Description  string    `gorm:"column:description"`
	UserID       uint      `gorm:"column:user_id"`
	Username     string    `gorm:"column:username"`
	CategoryID   uint      `gorm:"category_id"`
	CategoryName string    `gorm:"category_name"`
	Likes        uint      `gorm:"column:likes"`
	Answers      uint      `gorm:"column:answers"`
	UserLiked    uint      `gorm:"column:user_liked"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
