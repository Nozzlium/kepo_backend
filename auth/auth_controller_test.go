package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var userRepository = repository.UserRepositoryMock{Mock: mock.Mock{}}
var authService1 = AuthServiceImpl{UserRepository: &userRepository}
var authController = AuthControllerImpl{
	AuthService: &authService1,
	Validator:   validator.New(),
}

func TestPostRegisterSuccess(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Username: "user",
			Email:    "email.email@email.com",
		},
		nil,
	)
	user := requestbody.Register{
		Username: "user",
		Email:    "email.email@email.com",
		Password: "password",
	}
	userJson, _ := json.Marshal(&user)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/register", bytes.NewReader(userJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	authController.Register(recorder, request, nil)
	mockCall.Unset()

	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	body := response.UserWebResponse{}
	decoder.Decode(&body)

	assert.Equal(t, body.Code, http.StatusOK)

	assert.NotEqual(t, uint(0), body.Data.ID)
	assert.Equal(t, "user", body.Data.Username)
}

func TestInvalidEmailBody(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Username: "user",
			Email:    "email.email@email.com",
		},
		nil,
	)
	user := requestbody.Register{
		Username: "user",
		Email:    "supri",
		Password: "password",
	}
	userJson, _ := json.Marshal(&user)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/register", bytes.NewReader(userJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	assert.Panics(t, func() {
		authController.Register(recorder, request, nil)
	})
	mockCall.Unset()

}

func TestMissingParam(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Username: "user",
			Email:    "email.email@email.com",
		},
		nil,
	)
	user := requestbody.Register{
		Email:    "email.emai@email.com",
		Password: "password",
	}
	userJson, _ := json.Marshal(&user)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/register", bytes.NewReader(userJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	assert.Panics(t, func() {
		authController.Register(recorder, request, nil)
	})
	mockCall.Unset()

}

func TestServiceError(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{},
		gorm.ErrRecordNotFound,
	)
	user := requestbody.Register{
		Username: "user",
		Email:    "email.email@email.com",
		Password: "password",
	}
	userJson, _ := json.Marshal(&user)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/register", bytes.NewReader(userJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t, gorm.ErrRecordNotFound.Error(), func() {
		authController.Register(recorder, request, nil)
	})
	mockCall.Unset()
}

func TestPostLoginSuccess(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Username: "user",
			Email:    "email.email@email.com",
			Password: expectedHash,
		},
		nil,
	)
	cred := requestbody.Login{
		Identity: "user",
		Password: "password",
	}

	credJson, _ := json.Marshal(&cred)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/login", bytes.NewReader(credJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	authController.Login(recorder, request, nil)
	mockCall.Unset()

	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	body := response.AuthWebResponse{}
	decoder.Decode(&body)

	assert.Equal(t, body.Code, http.StatusOK)

	assert.NotNil(t, body.Data.Token)
	assert.NotEmpty(t, body.Data.Token)
}

func TestPostLoginNoUser(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{},
		gorm.ErrRecordNotFound,
	)
	cred := requestbody.Login{
		Identity: "supriadi",
		Password: "password",
	}

	credJson, _ := json.Marshal(&cred)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/login", bytes.NewReader(credJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t, constants.INVALID_CREDENTIAL, func() {
		authController.Login(recorder, request, nil)
	})
	mockCall.Unset()
}

func TestPostLoginInvalidPassword(t *testing.T) {
	mockCall := userRepository.Mock.On(
		"FindOneBasedOnIdentity",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.User{
			ID:       1,
			Username: "user",
			Email:    "email.email@email.com",
			Password: expectedHash,
		},
		nil,
	)
	cred := requestbody.Login{
		Identity: "user",
		Password: "supra",
	}

	credJson, _ := json.Marshal(&cred)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/login", bytes.NewReader(credJson))
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t, constants.INVALID_CREDENTIAL, func() {
		authController.Login(recorder, request, nil)
	})
	mockCall.Unset()
}
