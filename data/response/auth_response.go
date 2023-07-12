package response

type AuthResponse struct {
	Token string
}

type AuthWebResponse struct {
	BaseResponse
	Data AuthResponse `json:"data"`
}
