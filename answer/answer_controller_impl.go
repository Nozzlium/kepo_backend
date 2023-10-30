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

func NewAnswerController(
	answerService AnswerService,
	validator *validator.Validate,
) *AnswerControllerImpl {
	return &AnswerControllerImpl{
		AnswerService: answerService,
		Validator:     validator,
	}
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
	helper.WriteResponse(writer, &webResponse)
}

func (controller *AnswerControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userID uint = 0
	claims, err := tools.GetClaimsFromContext(request.Context())
	if err == nil {
		userID = claims.UserId
	}

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          userID,
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
	helper.WriteResponse(writer, &webResponse)
}

func (controller *AnswerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userID uint = 0
	claims, err := tools.GetClaimsFromContext(request.Context())
	if err == nil {
		userID = claims.UserId
	}

	idString := params.ByName("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	helper.PanicIfError(err)

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          userID,
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
	helper.WriteResponse(writer, &webResponse)

}

func (controller *AnswerControllerImpl) FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userID uint = 0
	claims, err := tools.GetClaimsFromContext(request.Context())
	if err == nil {
		userID = claims.UserId
	}

	questionUserIdString := params.ByName("id")
	questionUserId, err := strconv.ParseUint(questionUserIdString, 10, 32)
	if err != nil {
		panic(exception.BadRequestError{})
	}

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          userID,
		Answer: entity.Answer{
			UserID: uint(questionUserId),
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
	helper.WriteResponse(writer, &webResponse)
}

func (controller *AnswerControllerImpl) FindByQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userID uint = 0
	claims, err := tools.GetClaimsFromContext(request.Context())
	if err == nil {
		userID = claims.UserId
	}

	questionIdString := params.ByName("id")
	questionId, err := strconv.ParseUint(questionIdString, 10, 32)
	if err != nil {
		panic(exception.BadRequestError{})
	}

	answerParams := param.AnswerParam{
		PaginationParam: helper.GetPaginationParamFromQuerry(request),
		UserID:          userID,
		Answer: entity.Answer{
			QuestionID: uint(questionId),
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
	helper.WriteResponse(writer, &webResponse)
}

func (controller *AnswerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	idString := params.ByName("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	helper.PanicIfError(err)

	resp, err := controller.AnswerService.Delete(
		request.Context(),
		entity.Answer{
			ID: uint(id),
		},
	)
	helper.PanicIfError(err)

	webResponse := response.AnswerWebResponse{
		BaseResponse: response.BaseResponse{
			Code:   http.StatusOK,
			Status: constants.STATUS_OK,
		},
		Data: resp,
	}
	helper.WriteResponse(writer, &webResponse)
}

func (controller *AnswerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := tools.GetClaimsFromContext(request.Context())
	helper.PanicIfError(err)

	idString := params.ByName("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	helper.PanicIfError(err)

	body := requestbody.Answer{}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&body)
	helper.PanicIfError(err)

	resp, err := controller.AnswerService.Update(
		request.Context(),
		entity.Answer{
			ID:         uint(id),
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
		Data: resp,
	}
	helper.WriteResponse(writer, &webResponse)
}
