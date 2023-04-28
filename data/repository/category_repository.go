package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, category entity.Category) (entity.Category, error)
	FindAll(ctx context.Context, DB *gorm.DB) ([]entity.Category, error)
}
