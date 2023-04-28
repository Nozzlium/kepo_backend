package response

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
