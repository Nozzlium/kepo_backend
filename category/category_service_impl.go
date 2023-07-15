package category

import (
	"context"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DB                 *gorm.DB
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(
	DB *gorm.DB,
	categoryRepository repository.CategoryRepository,
) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB:                 DB,
		CategoryRepository: categoryRepository,
	}
}

func (service *CategoryServiceImpl) Get(ctx context.Context) ([]response.CategoryResponse, error) {
	categories, err := service.CategoryRepository.FindAll(
		ctx,
		service.DB,
	)
	return helper.CategoryEntityStructToResponseStruct(categories), err
}
