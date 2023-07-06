package exception

import (
	"encoding/json"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/response"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if badRequestError(writer, request, err) {
		return
	} else if invalidLoginError(writer, request, err) {
		return
	} else if notFoundError(writer, request, err) {
		return
	} else if unauthorizedError(writer, request, err) {
		return
	}

}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: constants.BAD_REQUEST,
			Data:   exception.Error(),
		}
		enc := json.NewEncoder(writer)
		enc.Encode(&webResponse)
	}
	return ok
}

func invalidLoginError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(InvalidLoginError)
	if ok {
		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: constants.INVALID_CREDENTIAL,
			Data:   exception.Error(),
		}
		enc := json.NewEncoder(writer)
		enc.Encode(&webResponse)
	}
	return ok
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: constants.NOT_FOUND,
			Data:   exception.Error(),
		}
		enc := json.NewEncoder(writer)
		enc.Encode(&webResponse)
	}
	return ok
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: constants.UNAUTHORIZED,
			Data:   exception.Error(),
		}
		enc := json.NewEncoder(writer)
		enc.Encode(&webResponse)
	}
	return ok
}
