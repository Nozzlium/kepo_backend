package response

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthWebResponse struct {
	BaseResponse
	Data AuthResponse `json:"data"`
}
