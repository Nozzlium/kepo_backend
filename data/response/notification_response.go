package response

type NotificationsWebResponse struct {
	BaseResponse
	Data NotificationsResponse `json:"data"`
}

type NotificationWebResponse struct {
	BaseResponse
	Data NotificationResponse `json:"data"`
}

type NotificationResponse struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"userId"`
	QuestionID       uint   `json:"questionId"`
	NotificationType string `json:"notificationType"`
	Headline         string `json:"headline"`
	Preview          string `json:"preview"`
	IsRead           bool   `json:"isRead"`
	CreatedAt        string `json:"createdAt"`
}

type NotificationsResponse struct {
	PageNo        int                    `json:"pageNo"`
	PageSize      int                    `json:"pageSize"`
	Notifications []NotificationResponse `json:"notifications"`
}
