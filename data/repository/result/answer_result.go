package result

type AnswerResult struct {
	ID         uint   `gorm:"column:id"`
	Content    string `gorm:"column:content"`
	QuestionID uint   `gorm:"column:question_id"`
	UserID     uint   `gorm:"column:user_id"`
	Username   string `gorm:"column:username"`
	Likes      uint   `gorm:"column:likes"`
	IsLiked    uint   `gorm:"column:user_liked"`
}
