package question

import (
	"encoding/json"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/customerror"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type QuestionControllerImpl struct {
	QuestionService QuestionService
	Validator       *validator.Validate
}

func (controller *QuestionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, ok := request.Context().Value(constants.USER_ID_CLAIMS).(tools.JwtClaims)
	if !ok {
		panic(customerror.UnauthorizedError{})
	}

	body := requestbody.Question{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	helper.PanicIfError(err)

	err = controller.Validator.Struct(body)
	if err != nil {
		panic(customerror.BadRequestError{})
	}

	question, err := controller.QuestionService.CreateQuestion(
		request.Context(),
		entity.Question{
			UserID:      claims.UserId,
			CategoryID:  body.CategoryID,
			Content:     body.Content,
			Description: body.Description,
		},
	)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constants.STATUS_OK,
		Data: response.QuestionCreate{
			ID:          question.ID,
			UserID:      question.UserID,
			CategoryID:  question.CategoryID,
			Content:     question.Content,
			Description: question.Description,
		},
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (controller *QuestionControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, ok := request.Context().Value(constants.USER_ID_CLAIMS).(tools.JwtClaims)
	if !ok {
		panic(customerror.UnauthorizedError{})
	}

	questionParam := param.InitQuestionParam()

	questionParam.PaginationParam = helper.GetPaginationParamFromQuerry(request)
	questionParam.UserID = claims.UserId

	questions, err := controller.QuestionService.FindAll(request.Context(), questionParam)
	helper.PanicIfError(err)

	questionsListResponse := response.QuestionsList{
		PageNo:    uint(questionParam.PageNo),
		PageSize:  uint(len(questions)),
		Questions: questions,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constants.STATUS_OK,
		Data:   questionsListResponse,
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (controller *QuestionControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, ok := request.Context().Value(constants.USER_ID_CLAIMS).(tools.JwtClaims)
	if !ok {
		panic(customerror.UnauthorizedError{})
	}

	questionParam := param.InitQuestionParam()
	questionParam.UserID = claims.UserId

	idString := params.ByName("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		panic(customerror.BadRequestError{})
	}
	questionParam.Question.ID = uint(id)

	questions, err := controller.QuestionService.FindAll(request.Context(), questionParam)
	helper.PanicIfError(err)

	if len(questions) < 1 {
		panic(customerror.NotFoundError{})
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constants.STATUS_OK,
		Data:   questions[0],
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (controller *QuestionControllerImpl) GetByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims, ok := request.Context().Value(constants.USER_ID_CLAIMS).(tools.JwtClaims)
	if !ok {
		panic(customerror.UnauthorizedError{})
	}

	userIdString := params.ByName("userId")
	userId, err := strconv.ParseUint(userIdString, 10, 32)
	if err != nil {
		panic(customerror.BadRequestError{})
	}

	questionParam := param.InitQuestionParam()
	questionParam.PaginationParam = helper.GetPaginationParamFromQuerry(request)
	questionParam.UserID = claims.UserId
	questionParam.Question.UserID = uint(userId)

	questions, err := controller.QuestionService.FindAll(
		request.Context(),
		questionParam,
	)
	helper.PanicIfError(err)

	questionsListResponse := response.QuestionsList{
		PageNo:    uint(questionParam.PageNo),
		PageSize:  uint(len(questions)),
		Questions: questions,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constants.STATUS_OK,
		Data:   questionsListResponse,
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}
