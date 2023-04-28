package entity

type Question struct {
	ID            uint
	UserID        uint
	CategoryID    uint
	Content       string
	Description   string
	QuestionLikes []QuestionLike
	Answers       []Answer
}
