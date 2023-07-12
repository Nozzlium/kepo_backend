package response

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserWebResponse struct {
	BaseResponse
	Data UserResponse `json:"data"`
}
