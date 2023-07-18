package auth

import (
	"context"
	"errors"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/mysqlerr"
	"nozzlium/kepo_backend/tools"

	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewAuthService(
	userRepository repository.UserRepository,
	DB *gorm.DB,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, param param.AuthParam) (response.UserResponse, error) {
	passHash, err := tools.HashPassword(param.Password)
	if err != nil {
		return response.UserResponse{}, err
	}
	user, err := service.UserRepository.Insert(
		ctx,
		service.DB,
		entity.User{
			Email:    param.Email,
			Username: param.Username,
			Password: passHash,
		},
	)
	err = mysqlerr.CheckMySQLError(err)
	if errors.Is(err, mysqlerr.DuplicateKeyError{}) {
		return response.UserResponse{}, exception.UserExistsError{}
	}
	return helper.UserEntityToResponse(user), err
}

func (service *AuthServiceImpl) Login(ctx context.Context, param param.LoginParam) (response.AuthResponse, error) {
	user, err := service.UserRepository.FindOneBasedOnIdentity(
		ctx,
		service.DB,
		entity.User{
			Username: param.Identity,
			Email:    param.Identity,
		},
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.AuthResponse{}, exception.InvalidLoginError{}
		}
		return response.AuthResponse{}, err
	}
	passCheck := tools.CheckPasswordHash(param.Password, user.Password)
	if !passCheck {
		return response.AuthResponse{}, exception.InvalidLoginError{}
	}
	token, err := tools.NewJwtToken(user.ID)
	return response.AuthResponse{Token: token}, err
}
