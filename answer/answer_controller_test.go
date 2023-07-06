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
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPostAnswerSuccess(t *testing.T) {
	mockCall := mockReturnCreatedAnswer()

	createJsonBytes, _ := json.Marshal(createAnswerBody)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2637/answer", bytes.NewReader(createJsonBytes))
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	answerController.Create(recorder, request.WithContext(ctx), httprouter.Params{})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.WebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)
	data := resp.Data.(map[string]interface{})

	assert.NotEqual(t, uint(0), data["id"])
	assert.Equal(t, tools.ClaimContext.UserId, uint(data["userId"].(float64)))
	assert.Equal(t, createAnswerBody.QuestionID, uint(data["questionId"].(float64)))
	assert.Equal(t, createAnswerBody.Content, data["content"])

	mockCall.Unset()
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
	resp := response.WebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)
	data := resp.Data.(map[string]interface{})

	assert.Equal(
		t,
		1,
		int(data["pageNo"].(float64)),
	)
	assert.Equal(
		t,
		len(expectedAnswers),
		int(data["pageSize"].(float64)),
	)

	answers := data["answers"].([]interface{})
	assert.Equal(
		t,
		len(answers),
		int(data["pageSize"].(float64)),
	)
	for i, answer := range answers {
		expectedAnswer := expectedAnswers[i]
		answerMap := answer.(map[string]interface{})
		assert.Equal(
			t,
			expectedAnswer.ID,
			uint(answerMap["id"].(float64)),
		)
		assert.Equal(
			t,
			expectedAnswer.QuestionID,
			uint(answerMap["questionId"].(float64)),
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
	resp := response.WebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)
	data := resp.Data.(map[string]interface{})

	assert.Equal(
		t,
		expectedAnswers[0].ID,
		uint(data["id"].(float64)),
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

	routerParam := httprouter.Param{
		Key:   "userId",
		Value: "1",
	}
	answerController.FindByUser(recorder, request.WithContext(ctx), httprouter.Params{
		routerParam,
	})

	res := recorder.Result()
	respBytes, _ := io.ReadAll(res.Body)
	resp := response.WebResponse{}
	json.Unmarshal(respBytes, &resp)

	assert.Equal(t, http.StatusOK, resp.Code)
	data := resp.Data.(map[string]interface{})

	assert.Equal(
		t,
		1,
		int(data["pageNo"].(float64)),
	)
	assert.Equal(
		t,
		len(expectedAnswers),
		int(data["pageSize"].(float64)),
	)

	answers := data["answers"].([]interface{})
	assert.Equal(
		t,
		len(answers),
		int(data["pageSize"].(float64)),
	)
	for i, answer := range answers {
		expectedAnswer := expectedAnswersFromSameUser[i]
		answerMap := answer.(map[string]interface{})
		assert.Equal(
			t,
			expectedAnswer.ID,
			uint(answerMap["id"].(float64)),
		)
		assert.Equal(
			t,
			expectedAnswer.QuestionID,
			uint(answerMap["questionId"].(float64)),
		)
		userMap := answerMap["user"].(map[string]interface{})
		userId := uint(userMap["id"].(float64))
		assert.Equal(
			t,
			expectedAnswer.UserID,
			userId,
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
