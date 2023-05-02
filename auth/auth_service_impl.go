package auth

import (
	"context"
	"errors"
	"nozzlium/kepo_backend/customerror"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"

	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func (service *AuthServiceImpl) Register(ctx context.Context, param param.AuthParam) (response.UserResponse, error) {
	user, err := service.UserRepository.Insert(
		ctx,
		service.DB,
		entity.User{
			Email:    param.Email,
			Username: param.Username,
			Password: param.Password,
		},
	)
	return helper.UserEntityToResponse(user), err
}

func (service *AuthServiceImpl) Login(ctx context.Context, param param.AuthParam) (response.AuthResponse, error) {
	user, err := service.UserRepository.FindOneBy(
		ctx,
		service.DB,
		entity.User{
			Username: param.Username,
			Email:    param.Email,
		},
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.AuthResponse{}, customerror.InvalidLoginError{}
		}
		return response.AuthResponse{}, err
	}
	passCheck := tools.CheckPasswordHash(param.Password, user.Password)
	if !passCheck {
		return response.AuthResponse{}, customerror.InvalidLoginError{}
	}
	token, err := tools.NewJwtToken(user.ID)
	return response.AuthResponse{Token: token}, err
}
