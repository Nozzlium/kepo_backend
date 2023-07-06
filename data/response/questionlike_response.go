package response

type QuestionLikeResponse struct {
	QuestionID uint `json:"questionId"`
	IsLiked    bool `json:"isLiked"`
	Likes      uint `json:"likes"`
}
