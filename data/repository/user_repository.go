package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error)
	FindOneBy(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error)
}
