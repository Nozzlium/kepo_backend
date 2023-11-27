package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type NotificationRepositoryMock struct {
}

func (mock *NotificationRepositoryMock) Create(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error) {
	panic("not implemented") // TODO: Implement
}

func (mock *NotificationRepositoryMock) FindBy(ctx context.Context, DB *gorm.DB, param param.NotificationParam) ([]entity.Notification, error) {
	panic("not implemented") // TODO: Implement
}

func (mock *NotificationRepositoryMock) Read(ctx context.Context, DB *gorm.DB, notification entity.Notification) (entity.Notification, error) {
	panic("not implemented") // TODO: Implement
}
