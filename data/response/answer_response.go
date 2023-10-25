package response

type AnswersWebResponse struct {
	BaseResponse
	Data AnswersResponse `json:"data"`
}

type AnswerWebResponse struct {
	BaseResponse
	Data AnswerResponse `json:"data"`
}

type AnswersResponse struct {
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
	Answers  []AnswerResponse `json:"answers"`
}

type AnswerResponse struct {
	ID         uint         `json:"id"`
	Content    string       `json:"content"`
	QuestionID uint         `json:"questionId"`
	User       UserResponse `json:"user"`
	Likes      uint         `json:"likes"`
	IsLiked    bool         `json:"isLiked"`
	CreatedAt  string       `json:"createdAt"`
}
