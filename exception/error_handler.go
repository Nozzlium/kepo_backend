package exception

import (
	"fmt"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	fmt.Println(err)
	if recordNotFound(writer, request, err) {
		return
	}
	if badRequestError(writer, request, err) {
		return
	}
	if invalidLoginError(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}
	if unauthorizedError(writer, request, err) {
		return
	}
	if userExistsError(writer, request, err) {
		return
	}

	webResponse := response.BaseResponse{
		Code:   http.StatusInternalServerError,
		Status: constants.INTERNAL_SERVER_ERROR,
	}
	helper.WriteResponse(writer, &webResponse)

}

func recordNotFound(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	ok := err == gorm.ErrRecordNotFound
	if ok {
		webResponse := response.BaseResponse{
			Code:   http.StatusNotFound,
			Status: constants.NOT_FOUND,
		}
		helper.WriteResponse(writer, &webResponse)
	}
	return ok
}

func userExistsError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UserExistsError)
	if ok {
		webResponse := response.BaseResponse{
			Code:   http.StatusConflict,
			Status: exception.Error(),
		}
		helper.WriteResponse(writer, &webResponse)
	}
	return ok
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		webResponse := response.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: exception.Error(),
		}
		helper.WriteResponse(writer, &webResponse)
	}
	return ok
}

func invalidLoginError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(InvalidLoginError)
	if ok {
		webResponse := response.BaseResponse{
			Code:   http.StatusUnauthorized,
			Status: exception.Error(),
		}
		helper.WriteResponse(writer, &webResponse)
	}
	return ok
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		webResponse := response.BaseResponse{
			Code:   http.StatusNotFound,
			Status: exception.Error(),
		}
		helper.WriteResponse(writer, &webResponse)
	}
	return ok
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		webResponse := response.BaseResponse{
			Code:   http.StatusUnauthorized,
			Status: exception.Error(),
		}
		helper.WriteResponse(writer, &webResponse)
	}
	return ok
}
