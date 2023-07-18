package requestbody

type QuestionLike struct {
	QuestionID uint `json:"questionId" validate:"required"`
	IsLike     bool `json:"isLike"`
}
