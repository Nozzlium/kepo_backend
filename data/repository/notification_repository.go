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
	Read(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error)
}
