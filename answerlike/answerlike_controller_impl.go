package answerlike

import (
	"encoding/json"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type AnswerLikeControllerImpl struct {
	AnswerLikeService AnswerLikeService
	Validator         *validator.Validate
}

func (controller *AnswerLikeControllerImpl) Like(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	body := requestbody.AnswerLike{}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&body)
	helper.PanicIfError(err)

	err = controller.Validator.Struct(body)
	helper.PanicIfError(err)

	resp, err := controller.AnswerLikeService.AssignLike(
		request.Context(),
		param.AnswerLikeParam{
			AnswerLike: entity.AnswerLike{
				AnswerID: body.AnswerID,
				UserID:   claims.UserId,
			},
			IsLike: body.IsLiked,
		},
	)
	helper.PanicIfError(err)

	webResp := response.WebResponse{
		Code:   http.StatusOK,
		Status: constants.STATUS_OK,
		Data:   resp,
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResp)
	helper.PanicIfError(err)
}
