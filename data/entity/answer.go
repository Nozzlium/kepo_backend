package entity

type Answer struct {
	ID          uint
	QuestionID  uint
	UserID      uint
	Content     string
	AnswerLikes []AnswerLike
}
