package response

type AnswerCreate struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"questionId"`
	UserID     uint   `json:"userId"`
	Content    string `json:"content"`
}
