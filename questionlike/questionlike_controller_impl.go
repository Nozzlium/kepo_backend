package questionlike

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

type QuestionLikeControllerImpl struct {
	QuestionLikeService QuestionLikeService
	Validator           *validator.Validate
}

func NewQuestionLikeController(
	questionLikeService QuestionLikeService,
	validator *validator.Validate,
) *QuestionLikeControllerImpl {
	return &QuestionLikeControllerImpl{
		QuestionLikeService: questionLikeService,
		Validator:           validator,
	}
}

func (controller *QuestionLikeControllerImpl) Like(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	body := requestbody.QuestionLike{}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&body)
	helper.PanicIfError(err)

	err = controller.Validator.Struct(body)
	helper.PanicIfError(err)

	resp, err := controller.QuestionLikeService.AssignLike(
		request.Context(),
		param.QuestionLikeParam{
			QuestionLike: entity.QuestionLike{
				QuestionID: body.QuestionID,
				UserID:     claims.UserId,
			},
			IsLiked: body.IsLike,
		},
	)
	helper.PanicIfError(err)

	webResponse := response.QuestionLikeWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp,
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(&webResponse)
}
