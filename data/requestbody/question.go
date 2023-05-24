package requestbody

type Question struct {
	CategoryID  uint   `json:"categoryId" validate:"required"`
	Content     string `json:"content" validate:"required"`
	Description string `json:"description" validate:"required"`
}
