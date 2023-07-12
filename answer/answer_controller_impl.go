package answer

import (
	"encoding/json"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type AnswerControllerImpl struct {
	AnswerService AnswerService
	Validator     *validator.Validate
}

func (controller *AnswerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	body := requestbody.Answer{}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&body)
	helper.PanicIfError(err)

	err = controller.Validator.Struct(body)
	if err != nil {
		panic(exception.BadRequestError{})
	}

	res, err := controller.AnswerService.CreateAnswer(
		request.Context(),
		entity.Answer{
			UserID:     claims.UserId,
			QuestionID: body.QuestionID,
			Content:    body.Content,
		},
	)
	helper.PanicIfError(err)

	webResponse := response.AnswerWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: res,
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (controller *AnswerControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          claims.UserId,
	}

	resp, err := controller.AnswerService.FindBy(request.Context(), answerParams)
	helper.PanicIfError(err)

	webResponse := response.AnswersWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: response.AnswersResponse{
			Page:     answerParams.PageNo,
			PageSize: len(resp),
			Answers:  resp,
		},
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (controller *AnswerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	idString := params.ByName("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	helper.PanicIfError(err)

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          claims.UserId,
		Answer: entity.Answer{
			ID: uint(id),
		},
	}

	resp, err := controller.AnswerService.FindBy(request.Context(), answerParams)
	helper.PanicIfError(err)

	if len(resp) < 1 {
		panic(exception.NotFoundError{})
	}

	webResponse := response.AnswerWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp[0],
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)

}

func (controller *AnswerControllerImpl) FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	userIdString := params.ByName("userId")
	userId, err := strconv.ParseUint(userIdString, 10, 32)
	if err != nil {
		panic(exception.BadRequestError{})
	}

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          claims.UserId,
		Answer: entity.Answer{
			UserID: uint(userId),
		},
	}

	resp, err := controller.AnswerService.FindBy(request.Context(), answerParams)
	helper.PanicIfError(err)

	webResponse := response.AnswersWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: response.AnswersResponse{
			Page:     answerParams.PageNo,
			PageSize: len(resp),
			Answers:  resp,
		},
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}
