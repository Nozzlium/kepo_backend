package response

type AnswerLikeWebResponse struct {
	BaseResponse
	Data AnswerLikeResponse `json:"data"`
}

type AnswerLikeResponse struct {
	IsLiked  bool `json:"isLiked"`
	AnswerID uint `json:"answerId"`
	Likes    uint `json:"likes"`
}
