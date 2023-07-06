package requestbody

type Answer struct {
	QuestionID uint   `json:"questionId" validate:"required,number"`
	Content    string `json:"content" validate:"required,max=300"`
}
