package question

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/tools"
	"strconv"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var questionController = QuestionControllerImpl{
	QuestionService: questionService,
	Validator:       validator.New(),
}

func TestPostCreateSuccess(t *testing.T) {
	mockCall := mockReturnInsertSuccess()
	mockCallGet := mockReturnOneDetailed()
	createJsonBytes, _ := json.Marshal(createRequestBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/question", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	questionController.Create(recorder, request.WithContext(ctx), nil)

	assert.Equal(t, http.StatusOK, recorder.Code)
	mockCall.Unset()
	mockCallGet.Unset()

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.QuestionWebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.NotEqual(t, uint(0), resp.Data.ID)
	assert.Equal(t, tools.ClaimContext.UserId, resp.Data.User.ID)
	assert.Equal(t, createRequestBody.CategoryID, uint(resp.Data.Category.ID))
}

func TestPostCreateError(t *testing.T) {
	mockCall := mockReturnInsertError()
	createJsonBytes, _ := json.Marshal(createRequestBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/question", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
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
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(t,
		exception.BadRequestError{}.Error(),
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

	assert.PanicsWithError(t, exception.UnauthorizedError{}.Error(), func() {
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

	webResponse := response.QuestionsWebResponse{}
	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)

	assert.Equal(t, 1, webResponse.Data.Page)
	assert.Equal(t, 2, webResponse.Data.PageSize)

	for i, question := range webResponse.Data.Questions {
		mockQuestion := expectedQuestions[i]
		assert.Equal(t, mockQuestion.ID, question.ID)
		assert.Equal(t, mockQuestion.UserID, question.User.ID)
		assert.Equal(t, mockQuestion.CategoryID, question.Category.ID)
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
		exception.UnauthorizedError{}.Error(),
		func() {
			questionController.Get(recorder, request, nil)
		})

	mockCall.Unset()
}

func TestControllerGetQuestionById(t *testing.T) {
	mockCall := mockReturnOneDetailed()

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

	webResponse := response.QuestionWebResponse{}
	res := recorder.Result()
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, expectedQuestions[0].ID, webResponse.Data.ID)
	assert.Equal(t, expectedQuestions[0].UserID, webResponse.Data.User.ID)
	assert.Equal(t, expectedQuestions[0].CategoryID, webResponse.Data.Category.ID)

	mockCall.Unset()
}

func TestGetQuestionByIdUnauthorized(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/question/1", nil)
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: "1",
	})
	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
		func() {
			questionController.GetById(
				recorder,
				request,
				params,
			)
		},
	)
}

func TestControllerGetQuestionByIdBadRequest(t *testing.T) {

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
		exception.BadRequestError{}.Error(),
		func() {
			questionController.GetById(
				recorder,
				request.WithContext(ctx),
				params,
			)
		},
	)
}

func TestGetOneQuestionNotFound(t *testing.T) {
	mockCall := mockReturnOneNotFound()
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
		gorm.ErrRecordNotFound.Error(),
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

	var expectedId uint = 1
	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: strconv.Itoa(int(expectedId)),
	})

	questionController.GetByUser(recorder, request.WithContext(ctx), params)

	webResponse := response.QuestionsWebResponse{}
	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, 1, webResponse.Data.Page)
	assert.Equal(t, 2, webResponse.Data.PageSize)

	for i, question := range webResponse.Data.Questions {
		mockQuestion := expectedQuestionsSameUser[i]
		assert.Equal(t, mockQuestion.ID, question.ID)
		assert.Equal(t, mockQuestion.UserID, question.User.ID)
		assert.Equal(t, expectedId, question.User.ID)
	}

	mockCall.Unset()
}

func TestGetQuestionsByUserUnauthorized(t *testing.T) {
	mockCall := mockReturnQuestionsSameUserSuccess()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/user/1/questions?pageNo=1&pageSize=2", nil)
	recorder := httptest.NewRecorder()

	params := httprouter.Params{}
	params = append(params, httprouter.Param{
		Key:   "id",
		Value: "1",
	})

	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
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
		Key:   "id",
		Value: "samsul",
	})

	assert.PanicsWithError(
		t,
		exception.BadRequestError{}.Error(),
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

	webResponse := response.QuestionsWebResponse{}
	resp := recorder.Result()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)

	assert.Equal(t, 1, webResponse.Data.Page)
	assert.Equal(t, 0, webResponse.Data.PageSize)

	assert.Empty(t, webResponse.Data.Questions)

	mockCall.Unset()
}
