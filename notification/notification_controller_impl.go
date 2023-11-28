package notification

import (
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"

	"github.com/julienschmidt/httprouter"
)

type NotificationControllerImpl struct {
	NotificationService NotificationService
}

func (controller *NotificationControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	param := param.NotificationParam{}
	param.PaginationParam = helper.GetPaginationParamFromQuerry(request)
	param.Notification = entity.Notification{
		UserID: claims.UserId,
	}

	resp, err := controller.NotificationService.Find(
		request.Context(),
		param,
	)
	helper.PanicIfError(err)

	webResponse := response.NotificationsWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp,
	}

	helper.WriteResponse(writer, &webResponse)
}

func (controller *NotificationControllerImpl) Read(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
