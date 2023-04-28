package response

type AnswerResponse struct {
	ID         uint         `json:"id"`
	Content    string       `json:"content"`
	QuestionID uint         `json:"questionId"`
	User       UserResponse `json:"user"`
	Likes      uint         `json:"likes"`
	IsLiked    bool         `json:"isLiked"`
}
