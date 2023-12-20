package notification

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type NotificationService interface {
	Find(ctx context.Context, param param.NotificationParam) (response.NotificationsResponse, error)
	Read(ctx context.Context, entity entity.Notification) (response.NotificationResponse, error)
	GetUnreadCount(ctx context.Context, entity entity.Notification) (int, error)
}
