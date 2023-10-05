package requestbody

type QuestionLike struct {
	QuestionID uint `json:"questionId" validate:"required"`
	IsLiked    bool `json:"isLiked"`
}
