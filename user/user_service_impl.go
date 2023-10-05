package user

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(
	userRepository repository.UserRepository,
	DB *gorm.DB,
) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *UserServiceImpl) FindOneBy(ctx context.Context, param param.UserParam) (response.UserResponse, error) {
	user, err := service.UserRepository.FindOneBy(ctx, service.DB, param.User)
	return helper.UserEntityToResponse(user), err
}
