package notification

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type NotificationServiceImpl struct {
	NotificationRepository repository.NotificationRepository
	DB                     *gorm.DB
}

func (service *NotificationServiceImpl) Find(ctx context.Context, param param.NotificationParam) (response.NotificationsResponse, error) {
	res, err := service.NotificationRepository.FindBy(ctx, service.DB, param)
	if err != nil {
		return response.NotificationsResponse{}, err
	}
	notifs := helper.NotificationEntitiesToResponses(res)
	unreadCount, err := service.NotificationRepository.GetUnreadCount(ctx, service.DB, param.Notification.UserID)
	if err != nil {
		return response.NotificationsResponse{}, err
	}
	return response.NotificationsResponse{
		Notifications: notifs,
		TotalUnread:   unreadCount,
		PageNo:        param.PageNo,
		PageSize:      len(notifs),
	}, nil
}

func (service *NotificationServiceImpl) Read(ctx context.Context, entity entity.Notification) (entity.Notification, error) {
	return service.NotificationRepository.Read(ctx, service.DB, entity)
}
