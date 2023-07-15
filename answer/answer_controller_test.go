package answer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/tools"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPostAnswerSuccess(t *testing.T) {
	mockCall := mockReturnCreatedAnswer()
	mockCallInserted := mockReturnOneAnswerDetailed()

	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	answerController.Create(recorder, request.WithContext(ctx), httprouter.Params{})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.AnswerWebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Equal(t, tools.ClaimContext.UserId, resp.Data.User.ID)
	assert.Equal(t, createAnswerBody.QuestionID, resp.Data.ID)
	assert.Equal(t, createAnswerBody.Content, resp.Data.Content)

	mockCall.Unset()
	mockCallInserted.Unset()
}

func TestPostUnauthorized(t *testing.T) {
	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer", bytes.NewReader(createJsonBytes))
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
		func() {
			answerController.Create(recorder, request, httprouter.Params{})
		},
	)
}

func TestPostAnswerMissingParamInBody(t *testing.T) {
	mockCall := mockReturnCreatedAnswer()

	createJsonBytes, _ := json.Marshal(requestbody.Answer{
		Content: "test",
	})
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.BadRequestError{}.Error(),
		func() {
			answerController.Create(recorder, request.WithContext(ctx), httprouter.Params{})
		},
	)
	mockCall.Unset()
}

func TestPostAnswerExceedMaxLength(t *testing.T) {
	createJsonBytes, _ := json.Marshal(createAnswerBodyExceedLength)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.BadRequestError{}.Error(),
		func() {
			answerController.Create(recorder, request.WithContext(ctx), httprouter.Params{})
		},
	)
}

func TestServiceError(t *testing.T) {
	mockCall := mockReturnAnswersError()

	createJsonBytes, _ := json.Marshal(createAnswerBodyExceedLength)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.Panics(
		t,
		func() {
			answerController.Create(recorder, request.WithContext(ctx), httprouter.Params{})
		},
	)

	mockCall.Unset()
}

func TestGetQuestionsOK(t *testing.T) {
	mockCall := mockReturnAnswers()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer?pageNo=1&pageSize=5", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	answerController.Find(recorder, request.WithContext(ctx), httprouter.Params{})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.AnswersWebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Equal(
		t,
		1,
		resp.Data.Page,
	)
	assert.Equal(
		t,
		len(expectedAnswers),
		resp.Data.PageSize,
	)
	for i, answer := range resp.Data.Answers {
		expectedAnswer := expectedAnswers[i]
		assert.Equal(
			t,
			expectedAnswer.ID,
			answer.ID,
		)
		assert.Equal(
			t,
			expectedAnswer.QuestionID,
			answer.QuestionID,
		)
	}

	mockCall.Unset()
}

func TestGetAnswersUnauthorized(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer?pageNo=1&pageSize=5", nil)
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
		func() {
			answerController.Find(recorder, request, httprouter.Params{})
		},
	)
}

func TestGetAnswerServiceError(t *testing.T) {

	mockCall := mockReturnAnswersError()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer?pageNo=1&pageSize=5", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.Panics(
		t,
		func() {
			answerController.Find(recorder, request.WithContext(ctx), httprouter.Params{})
		},
	)

	mockCall.Unset()

}

func TestGetAnswersById(t *testing.T) {
	mockCall := mockReturnOneAnswer()

	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer/1", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	answerController.FindById(recorder, request.WithContext(ctx), httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: "1",
		},
	})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.QuestionWebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Equal(
		t,
		expectedAnswers[0].ID,
		resp.Data.ID,
	)

	mockCall.Unset()
}

func TestGetAnswerByIdUnauthorized(t *testing.T) {
	mockCall := mockReturnOneAnswer()

	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer/1", bytes.NewReader(createJsonBytes))
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
		func() {
			answerController.FindById(recorder, request, httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: "1",
				},
			})
		},
	)

	mockCall.Unset()
}

func TestGetAnswerByIdServiceError(t *testing.T) {
	mockCall := mockReturnAnswersError()

	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer/1", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.Panics(
		t,
		func() {
			answerController.FindById(recorder, request.WithContext(ctx), httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: "1",
				},
			})
		},
	)

	mockCall.Unset()
}

func TestGetAnswerNotFound(t *testing.T) {
	mockCall := mockReturnEmpty()

	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer/1", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.NotFoundError{}.Error(),
		func() {
			answerController.FindById(recorder, request.WithContext(ctx), httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: "1",
				},
			})
		},
	)

	mockCall.Unset()
}

func TestGetAnswerByUser(t *testing.T) {
	mockCall := mockReturnAnswersFromSameUser()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/users/1/answer?pageNo=1&pageSize=10", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	var userId uint = 1
	routerParam := httprouter.Param{
		Key:   "userId",
		Value: strconv.FormatUint(uint64(userId), 10),
	}
	answerController.FindByUser(recorder, request.WithContext(ctx), httprouter.Params{
		routerParam,
	})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.AnswersWebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Equal(
		t,
		1,
		resp.Data.Page,
	)
	assert.Equal(
		t,
		len(expectedAnswers),
		resp.Data.PageSize,
	)

	for i, answer := range resp.Data.Answers {
		expectedAnswer := expectedAnswersFromSameUser[i]
		assert.Equal(
			t,
			expectedAnswer.ID,
			answer.ID,
		)
		assert.Equal(
			t,
			expectedAnswer.QuestionID,
			answer.QuestionID,
		)
		assert.Equal(
			t,
			userId,
			answer.User.ID,
		)
		assert.Equal(
			t,
			expectedAnswer.UserID,
			answer.User.ID,
		)
	}

	mockCall.Unset()
}

func TestGetAnswersByUserUnauthorized(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/users/1/answer?pageNo=1&pageSize=10", nil)
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
		func() {
			answerController.FindByUser(recorder, request, httprouter.Params{
				httprouter.Param{
					Key:   "userId",
					Value: "1",
				},
			})
		},
	)
}

func TestGetAnswerByUserInvalidUserId(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/users/1/answer?pageNo=1&pageSize=10", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.BadRequestError{}.Error(),
		func() {
			answerController.FindByUser(recorder, request.WithContext(ctx), httprouter.Params{
				httprouter.Param{
					Key:   "userId",
					Value: "samsul",
				},
			})
		},
	)
}

func TestGetAnswerByUserServiceError(t *testing.T) {
	mockCall := mockReturnAnswersError()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/users/1/answer?pageNo=1&pageSize=10", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	assert.Panics(
		t,
		func() {
			answerController.FindByUser(recorder, request.WithContext(ctx), httprouter.Params{
				httprouter.Param{
					Key:   "userId",
					Value: "1",
				},
			})
		},
	)

	mockCall.Unset()
}

func TestGetAnswerByQuestion(t *testing.T) {
	mockCall := mockReturnAnswerByQuestion()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/users/1/answer?pageNo=1&pageSize=10", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	var questionId uint = 1
	routerParam := httprouter.Param{
		Key:   "questionId",
		Value: strconv.FormatUint(uint64(questionId), 10),
	}
	answerController.FindByQuestion(recorder, request.WithContext(ctx), httprouter.Params{
		routerParam,
	})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.AnswersWebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Equal(
		t,
		1,
		resp.Data.Page,
	)
	assert.Equal(
		t,
		len(expectedAnswers),
		resp.Data.PageSize,
	)

	for i, answer := range resp.Data.Answers {
		expectedAnswer := expectedAnswersFromSameUser[i]
		assert.Equal(
			t,
			expectedAnswer.ID,
			answer.ID,
		)
		assert.Equal(
			t,
			expectedAnswer.QuestionID,
			answer.QuestionID,
		)
		assert.Equal(
			t,
			questionId,
			answer.User.ID,
		)
		assert.Equal(
			t,
			expectedAnswer.UserID,
			answer.User.ID,
		)
	}

	mockCall.Unset()
}
