package requestbody

type AnswerLike struct {
	AnswerID uint `json:"answerId" validate:"required,number"`
	IsLiked  bool `json:"isLiked"`
}
