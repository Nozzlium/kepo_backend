package response

type NotificationCountWebResponse struct {
	BaseResponse
	Data NotificationCountResponse `json:"data"`
}

type NotificationCountResponse struct {
	TotalUnread int `json:"totalUnread"`
}
