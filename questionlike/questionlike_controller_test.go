package questionlike

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikeQuestionSuccess(t *testing.T) {

	mockCallQuestionLike := mockLikeQuestion()
	mockCallQuestion := mockGetLikedAnswer()

	bodyBytes, _ := json.Marshal(&likeRequestBody)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:2637/question_like",
		bytes.NewReader(bodyBytes),
	)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	questionLikeController.Like(
		recorder,
		request.WithContext(ctx),
		nil,
	)

	webResponse := response.QuestionLikeWebResponse{}
	resultBytes, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(resultBytes, &webResponse)

	assert.Equal(
		t,
		expectedLikedQuestion[0].ID,
		webResponse.Data.QuestionID,
	)
	assert.Equal(
		t,
		true,
		webResponse.Data.IsLiked,
	)

	mockCallQuestion.Unset()
	mockCallQuestionLike.Unset()

}

func TestDislikeQuestionSuccess(t *testing.T) {

	mockCallQuestionLike := mockDislikeQuestion()
	mockCallQuestion := mockGetDislikedAnswer()

	bodyBytes, _ := json.Marshal(&dislikeRequestBody)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:2637/question_like",
		bytes.NewReader(bodyBytes),
	)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	questionLikeController.Like(
		recorder,
		request.WithContext(ctx),
		nil,
	)

	webResponse := response.QuestionLikeWebResponse{}
	resultBytes, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(resultBytes, &webResponse)

	assert.Equal(
		t,
		expectedDislikedQuestion[0].ID,
		webResponse.Data.QuestionID,
	)
	assert.Equal(
		t,
		false,
		webResponse.Data.IsLiked,
	)

	mockCallQuestion.Unset()
	mockCallQuestionLike.Unset()

}

func TestLikeQuestionUnauthorized(t *testing.T) {
	bodyBytes, _ := json.Marshal(&dislikeRequestBody)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:2637/question_like",
		bytes.NewReader(bodyBytes),
	)
	recorder := httptest.NewRecorder()

	assert.PanicsWithError(
		t,
		exception.UnauthorizedError{}.Error(),
		func() {
			questionLikeController.Like(
				recorder,
				request,
				nil,
			)
		},
	)

}
