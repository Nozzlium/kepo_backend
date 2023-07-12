package response

type QuestionWebResponse struct {
	BaseResponse
	Data QuestionResponse `json:"data"`
}

type QuestionsWebResponse struct {
	BaseResponse
	Data QuestionsResponse `json:"data"`
}

type QuestionsResponse struct {
	Page      int                `json:"page"`
	PageSize  int                `json:"pageSize"`
	Questions []QuestionResponse `json:"questions"`
}

type QuestionResponse struct {
	ID          uint             `json:"id"`
	User        UserResponse     `json:"user"`
	Category    CategoryResponse `json:"category"`
	Content     string           `json:"content"`
	Description string           `json:"description"`
	Likes       uint             `json:"likes"`
	Answers     uint             `json:"answers"`
	IsLiked     bool             `json:"isLiked"`
}
