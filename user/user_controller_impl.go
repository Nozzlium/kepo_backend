package user

import (
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService UserService
}

func NewUserController(
	userService UserService,
) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) GetDetails(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	userParam := param.UserParam{}
	userParam.User = entity.User{
		ID: uint(claims.UserId),
	}

	user, err := controller.UserService.FindOneBy(request.Context(), userParam)
	helper.PanicIfError(err)

	userResponse := response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
	webResponse := response.UserWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: userResponse,
	}

	helper.WriteResponse(writer, &webResponse)
}

func (controller *UserControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userParam := param.UserParam{}

	idString := params.ByName("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		panic(exception.BadRequestError{})
	}
	userParam.User = entity.User{
		ID: uint(id),
	}

	user, err := controller.UserService.FindOneBy(request.Context(), userParam)
	helper.PanicIfError(err)

	userResponse := response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
	webResponse := response.UserWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: userResponse,
	}

	helper.WriteResponse(writer, &webResponse)
}
