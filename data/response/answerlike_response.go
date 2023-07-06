package response

type AnswerLikeResponse struct {
	IsLiked  bool `json:"isLiked"`
	AnswerID uint `json:"answerId"`
	Likes    uint `json:"likes"`
}
