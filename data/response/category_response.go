package response

type CategoryWebResponse struct {
	BaseResponse
	Data []CategoryResponse `json:"data"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
