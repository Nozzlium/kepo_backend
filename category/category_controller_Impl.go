package category

import (
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService CategoryService
}

func NewCategoryController(
	categoryService CategoryService,
) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	resp, err := controller.CategoryService.Get(request.Context())
	helper.PanicIfError(err)

	webResponse := response.CategoryWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp,
	}
	helper.WriteResponse(writer, &webResponse)
}
