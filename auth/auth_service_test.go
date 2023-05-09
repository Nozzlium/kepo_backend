package auth

import (
	"context"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/customerror"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/repositorymock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var expectedHash = "$2a$10$A.fc97ZQ0O3MBQuaA4F5L.PATPV3tbOS.Fli.5nF79eVDTSVFt0xW"

var userRepositoryMock = repositorymock.UserRepositoryMock{Mock: mock.Mock{}}
var authService = AuthServiceImpl{
	UserRepository: &userRepositoryMock,
}

var expectedUser = entity.User{
	ID:       1,
	Email:    "email.email@email.com",
	Username: "user",
}

func TestRegisterSuccess(t *testing.T) {
	mockCall := userRepositoryMock.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Email:    "email.email@email.com",
			Username: "user",
		},
		nil,
	)

	param := param.AuthParam{
		Email:    "email.email@email.com",
		Username: "user",
		Password: "password",
	}
	response, err := authService.Register(
		context.Background(),
		param,
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Username, response.Username)
	assert.Equal(t, expectedUser.ID, response.ID)
}

func TestRegisterError(t *testing.T) {
	mockCall := userRepositoryMock.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{},
		gorm.ErrInvalidData,
	)
	_, err := authService.Register(
		context.Background(),
		param.AuthParam{
			Username: "user",
			Email:    "email.email@email.com",
			Password: "password",
		},
	)
	mockCall.Unset()

	assert.NotNil(t, err)
}

func TestLoginSuccess(t *testing.T) {
	mockCall := userRepositoryMock.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Email:    "email.email@email.com",
			Username: "user",
			Password: expectedHash,
		},
		nil,
	)
	response, err := authService.Login(
		context.Background(),
		param.LoginParam{
			Identity: "user",
			Password: "password",
		},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.NotNil(t, response.Token)
	assert.NotEmpty(t, response.Token)
}

func TestLoginError(t *testing.T) {
	mockCall := userRepositoryMock.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{},
		gorm.ErrUnsupportedDriver,
	)
	_, err := authService.Login(
		context.Background(),
		param.LoginParam{
			Identity: "wadidawwww",
			Password: "mukeluwadidaw",
		},
	)
	mockCall.Unset()

	assert.NotNil(t, err)
}

func TestLoginNoUser(t *testing.T) {
	mockCall := userRepositoryMock.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{},
		gorm.ErrRecordNotFound,
	)
	_, err := authService.Login(
		context.Background(),
		param.LoginParam{
			Identity: "wadidawwww",
			Password: "mukeluwadidaw",
		},
	)
	mockCall.Unset()

	assert.NotNil(t, err)
	assert.IsType(t, customerror.InvalidLoginError{}, err)
	assert.Equal(t, constants.INVALID_CREDENTIAL, err.Error())
}

func TestLoginInvalidPassword(t *testing.T) {
	mockCall := userRepositoryMock.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Email:    "email.email@email.com",
			Username: "user",
			Password: expectedHash,
		},
		nil,
	)
	_, err := authService.Login(
		context.Background(),
		param.LoginParam{
			Identity: "user",
			Password: "kicoetcoet",
		},
	)
	mockCall.Unset()

	assert.NotNil(t, err)
	assert.IsType(t, customerror.InvalidLoginError{}, err)
	assert.Equal(t, constants.INVALID_CREDENTIAL, err.Error())
}
