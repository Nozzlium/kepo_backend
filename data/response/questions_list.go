package response

type QuestionsList struct {
	Questions []QuestionResponse `json:"questions"`
	PageNo    uint               `json:"pageNo"`
	PageSize  uint               `json:"pageSize"`
}
