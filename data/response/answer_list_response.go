package response

type AnswerListResponse struct {
	PageNo   uint             `json:"pageNo"`
	PageSize uint             `json:"pageSize"`
	Answers  []AnswerResponse `json:"answers"`
}
