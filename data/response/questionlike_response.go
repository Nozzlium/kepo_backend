package response

type QuestionLikeWebResponse struct {
	BaseResponse
	Data QuestionLikeResponse `json:"data"`
}

type QuestionLikeResponse struct {
	QuestionID uint `json:"questionId"`
	IsLiked    bool `json:"isLiked"`
	Likes      uint `json:"likes"`
}
