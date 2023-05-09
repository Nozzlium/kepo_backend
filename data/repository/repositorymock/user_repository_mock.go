package repositorymock

import (
	"context"
	"nozzlium/kepo_backend/data/entity"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) Insert(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		user,
	)
	args0 := args[0].(entity.User)
	args1 := args[1]
	if args1 == nil {
		return args0, nil
	}
	return args0, args1.(error)
}

func (repository *UserRepositoryMock) FindOneBy(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		user,
	)
	args0 := args[0].(entity.User)
	args1 := args[1]
	if args1 == nil {
		return args0, nil
	}
	return args0, args1.(error)
}

func (repository *UserRepositoryMock) FindOneBasedOnIdentity(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		user,
	)
	args0 := args[0].(entity.User)
	args1 := args[1]
	if args1 == nil {
		return args0, nil
	}
	return args0, args1.(error)
}
