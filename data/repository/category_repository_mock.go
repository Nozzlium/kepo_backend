package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type CategoryRepositoryMock struct {
	Mock *mock.Mock
}

func (repository *CategoryRepositoryMock) Insert(ctx context.Context, DB *gorm.DB, category entity.Category) (entity.Category, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryMock) FindAll(ctx context.Context, DB *gorm.DB) ([]entity.Category, error) {
	args := repository.Mock.Called(ctx, DB)
	args0 := args.Get(0)
	args1 := args.Get(1)
	if args1 == nil {
		return args0.([]entity.Category), nil
	}
	return args0.([]entity.Category), args1.(error)
}
