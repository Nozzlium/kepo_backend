package response

type QuestionCreate struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"userId"`
	CategoryID  uint   `json:"categoryId"`
	Content     string `json:"content"`
	Description string `json:"description"`
}
