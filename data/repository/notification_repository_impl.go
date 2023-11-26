package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type NotificationRepositoryImpl struct {
}

func (repository *NotificationRepositoryImpl) Create(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error) {
	insert := DB.WithContext(ctx).Create(&notification)
	return notification, insert.Error
}

func (repository *NotificationRepositoryImpl) FindBy(ctx context.Context, DB *gorm.DB, param param.NotificationParam) ([]entity.Notification, error) {
	notifications := []entity.Notification{}
	find := DB.WithContext(ctx).
		Where("user_id = ?", param.Notification.UserID).
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Find(&notifications)
	return notifications, find.Error
}

func (repository *NotificationRepositoryImpl) Read(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error) {
	update := DB.WithContext(ctx).
		Model(&notification).
		Where("id = ?", notification.ID).
		Update("is_read", true)
	return notification, update.Error
}
