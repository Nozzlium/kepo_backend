package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService AuthService
	Validator   *validator.Validate
}

func NewAuthController(
	authService AuthService,
	validator *validator.Validate,
) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
		Validator:   validator,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	body := requestbody.Register{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	helper.PanicIfError(err)

	err = controller.Validator.Struct(body)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println("har", validationErrors[0].Field())
		helper.PanicIfError(err)
	}

	resp, err := controller.AuthService.Register(
		request.Context(),
		param.AuthParam{
			Username: body.Username,
			Email:    body.Email,
			Password: body.Password,
		},
	)
	helper.PanicIfError(err)

	respBody := response.UserWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp,
	}
	helper.WriteResponse(writer, &respBody)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	body := requestbody.Login{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	helper.PanicIfError(err)

	err = controller.Validator.Struct(body)
	helper.PanicIfError(err)

	param := param.LoginParam{
		Identity: body.Identity,
		Password: body.Password,
	}
	resp, err := controller.AuthService.Login(
		request.Context(),
		param,
	)
	helper.PanicIfError(err)

	respBody := response.AuthWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp,
	}
	helper.WriteResponse(writer, &respBody)
}
