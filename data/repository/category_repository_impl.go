package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, category entity.Category) (entity.Category, error) {
	insert := DB.WithContext(ctx).Create(&category)
	return category, insert.Error
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, DB *gorm.DB) ([]entity.Category, error) {
	categories := []entity.Category{}
	find := DB.WithContext(ctx).Find(&categories)
	return categories, find.Error
}
