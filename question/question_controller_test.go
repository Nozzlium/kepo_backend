package question

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/customerror"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/tools"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var questionController = QuestionControllerImpl{
	QuestionService: questionService,
	Validator:       validator.New(),
}

func TestPostCreateSuccess(t *testing.T) {
	mockCall := mockReturnInsertSuccess()
	createJsonBytes, _ := json.Marshal(createRequestBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/question", bytes.NewReader(createJsonBytes))
	ctx := getClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	questionController.Create(recorder, request.WithContext(ctx), nil)

	assert.Equal(t, http.StatusOK, recorder.Code)
	mockCall.Unset()

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.WebResponse{}
	json.Unmarshal(respBytes, &resp)
	data := resp.Data.(map[string]interface{})

	assert.NotEqual(t, uint(0), data["id"])
	assert.Equal(t, claimContext.UserId, uint(data["userId"].(float64)))
	assert.Equal(t, createRequestBody.CategoryID, uint(data["categoryId"].(float64)))
}

func TestPostCreateError(t *testing.T) {
	mockCall := mockReturnInsertError()
	createJsonBytes, _ := json.Marshal(createRequestBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/question", bytes.NewReader(createJsonBytes))
	ctx := getClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.Panics(t, func() {
		questionController.Create(recorder, request.WithContext(ctx), nil)
	})
	mockCall.Unset()
}

func TestPostCreateBodyMissingParam(t *testing.T) {
	mockCall := mockReturnInsertError()
	createJsonBytes, _ := json.Marshal(requestbody.Question{
		Content:     "haha",
		Description: "hihi",
	})
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/question", bytes.NewReader(createJsonBytes))
	ctx := getClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t,
		customerror.BadRequestError{}.Error(),
		func() {
			questionController.Create(recorder, request.WithContext(ctx), nil)
		})
	mockCall.Unset()
}

func TestPostCreateNoClaims(t *testing.T) {
	mockCall := mockReturnInsertSuccess()
	createJsonBytes, _ := json.Marshal(createRequestBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/question", bytes.NewReader(createJsonBytes))
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t, customerror.UnauthorizedError{}.Error(), func() {
		questionController.Create(recorder, request, nil)
	})
	mockCall.Unset()
}

func TestControllerGetQuestionsSuccess(t *testing.T) {
	mockCall := mockReturnQuestionsSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question?pageNo=1&pageSize=10", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	questionController.Get(recorder, request.WithContext(ctx), nil)

	webResponse := response.WebResponse{}
	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)

	dataResponse := webResponse.Data.(map[string]interface{})
	pageNo := uint(dataResponse["pageNo"].(float64))
	pageSize := uint(dataResponse["pageSize"].(float64))
	assert.Equal(t, uint(1), pageNo)
	assert.Equal(t, uint(2), pageSize)

	questions := dataResponse["questions"].([]interface{})
	for i, question := range questions {
		mockQuestion := expectedQuestions[i]
		questionId := uint(question.(map[string]interface{})["id"].(float64))
		assert.Equal(t, mockQuestion.ID, questionId)
	}

	mockCall.Unset()
}

func TestControllerGetQuestionServiceError(t *testing.T) {
	mockCall := mockReturnQuestionError()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question?pageNo=1&pageSize=10", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	assert.Panics(t, func() {
		questionController.Get(recorder, request.WithContext(ctx), nil)
	})

	mockCall.Unset()

}

func TestGetQuestionsUnauthorized(t *testing.T) {
	mockCall := mockReturnQuestionError()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question?pageNo=1&pageSize=10", nil)
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t,
		customerror.UnauthorizedError{}.Error(),
		func() {
			questionController.Get(recorder, request, nil)
		})

	mockCall.Unset()
}

func TestControllerGetQuestionById(t *testing.T) {
	mockCall := mockReturnOneQuestionSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question/1", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: "1",
	})
	questionController.GetById(
		recorder,
		request.WithContext(ctx),
		params,
	)

	webResponse := response.WebResponse{}
	res := recorder.Result()
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)

	dataResponse := webResponse.Data.(map[string]interface{})
	assert.Equal(t, expectedOneQuestion[0].ID, uint(dataResponse["id"].(float64)))

	mockCall.Unset()
}

func TestGetQuestionByIdUnauthorized(t *testing.T) {
	mockCall := mockReturnOneQuestionSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question/1", nil)
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: "1",
	})
	assert.PanicsWithError(
		t,
		customerror.UnauthorizedError{}.Error(),
		func() {
			questionController.GetById(
				recorder,
				request,
				params,
			)
		},
	)

	mockCall.Unset()
}

func TestControllerGetQuestionByIdBadRequest(t *testing.T) {
	mockCall := mockReturnOneQuestionSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question/1", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: "samsul",
	})
	assert.PanicsWithError(
		t,
		customerror.BadRequestError{}.Error(),
		func() {
			questionController.GetById(
				recorder,
				request.WithContext(ctx),
				params,
			)
		},
	)

	mockCall.Unset()
}

func TestGetOneQuestionNotFound(t *testing.T) {
	mockCall := mockReturnEmptyQuestion()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question/1", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: "1",
	})
	assert.PanicsWithError(
		t,
		customerror.NotFoundError{}.Error(),
		func() {
			questionController.GetById(
				recorder,
				request.WithContext(ctx),
				params,
			)
		},
	)
	mockCall.Unset()
}

func TestGetQuestionsByUserSuccess(t *testing.T) {
	mockCall := mockReturnQuestionsSameUserSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/user/1/questions?pageNo=1&pageSize=2", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "userId",
		Value: "1",
	})

	questionController.GetByUser(recorder, request.WithContext(ctx), params)

	webResponse := response.WebResponse{}
	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)

	dataResponse := webResponse.Data.(map[string]interface{})
	pageNo := uint(dataResponse["pageNo"].(float64))
	pageSize := uint(dataResponse["pageSize"].(float64))
	assert.Equal(t, uint(1), pageNo)
	assert.Equal(t, uint(2), pageSize)

	questions := dataResponse["questions"].([]interface{})
	for i, question := range questions {
		mockQuestion := expectedQuestionsSameUser[i]
		questionId := uint(question.(map[string]interface{})["id"].(float64))
		assert.Equal(t, mockQuestion.ID, questionId)
		assert.Equal(t, uint(1), mockQuestion.UserID)
	}

	mockCall.Unset()
}

func TestGetQuestionsByUserUnauthorized(t *testing.T) {
	mockCall := mockReturnQuestionsSameUserSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/user/1/questions?pageNo=1&pageSize=2", nil)
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "userId",
		Value: "1",
	})

	assert.PanicsWithError(
		t,
		customerror.UnauthorizedError{}.Error(),
		func() {
			questionController.GetByUser(recorder, request, params)
		},
	)

	mockCall.Unset()
}

func TestGetQuestionsByUserInvalidUserId(t *testing.T) {
	mockCall := mockReturnQuestionsSameUserSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/user/1/questions?pageNo=1&pageSize=2", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "userId",
		Value: "samsul",
	})

	assert.PanicsWithError(
		t,
		customerror.BadRequestError{}.Error(),
		func() {
			questionController.GetByUser(recorder, request.WithContext(ctx), params)
		},
	)

	mockCall.Unset()
}

func TestGetQuestionNoData(t *testing.T) {
	mockCall := mockReturnEmptyQuestion()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question?pageNo=1&pageSize=10", nil)
	ctx := context.WithValue(request.Context(), constants.USER_ID_CLAIMS, tools.JwtClaims{
		UserId: 1,
	})
	recorder := httptest.NewRecorder()

	questionController.Get(recorder, request.WithContext(ctx), nil)

	webResponse := response.WebResponse{}
	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)

	dataResponse := webResponse.Data.(map[string]interface{})
	pageNo := uint(dataResponse["pageNo"].(float64))
	pageSize := uint(dataResponse["pageSize"].(float64))
	assert.Equal(t, uint(1), pageNo)
	assert.Equal(t, uint(0), pageSize)

	questions := dataResponse["questions"].([]interface{})
	assert.Empty(t, questions)

	mockCall.Unset()
}
