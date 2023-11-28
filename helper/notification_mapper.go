package helper

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/response"
)

func NotificationEntityToResponse(entity entity.Notification) response.NotificationResponse {
	return response.NotificationResponse{
		ID:               entity.ID,
		UserID:           entity.UserID,
		QuestionID:       entity.QuestionID,
		NotificationType: entity.NotifType,
		Headline:         entity.Headline,
		Preview:          entity.Preview,
		IsRead:           entity.IsRead,
		CreatedAt:        TimeToString(entity.CreatedAt),
	}
}

func NotificationEntitiesToResponses(entities []entity.Notification) []response.NotificationResponse {
	responses := []response.NotificationResponse{}
	for _, entity := range entities {
		responses = append(responses, NotificationEntityToResponse(entity))
	}
	return responses
}
