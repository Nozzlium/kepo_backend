package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error)
	FindBy(ctx context.Context, DB *gorm.DB, param param.NotificationParam) ([]entity.Notification, error)
	FindOneBy(ctx context.Context, DB *gorm.DB, entity entity.Notification) (entity.Notification, error)
	Read(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error)
	GetUnreadCount(ctx context.Context, DB *gorm.DB, userId uint) (int, error)
}
