package answerlike

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/tools"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestLikeSuccess(t *testing.T) {

	mockLike := mockInsertLike()
	mockGetAnswer := mockGetLikedAnswer()

	bodyBytes, _ := json.Marshal(&likeRequestBody)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:2637/question_like",
		bytes.NewReader(bodyBytes),
	)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	answerLikeController.Like(
		recorder,
		request.WithContext(ctx),
		httprouter.Params{},
	)

	webResp := response.WebResponse{}
	dec := json.NewDecoder(recorder.Result().Body)
	dec.Decode(&webResp)

	assert.Equal(t, http.StatusOK, webResp.Code)

	data := webResp.Data.(map[string]interface{})
	assert.Equal(
		t,
		likedAnswer[0].ID,
		uint(data["answerId"].(float64)),
	)
	assert.Equal(
		t,
		true,
		data["isLiked"].(bool),
	)

	mockLike.Unset()
	mockGetAnswer.Unset()

}

func TestDislikeSuccess(t *testing.T) {

	mockLike := mockRemoveLike()
	mockGetAnswer := mockGetDislikedAnswer()

	bodyBytes, _ := json.Marshal(&dislikeRequestBody)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:2637/question_like",
		bytes.NewReader(bodyBytes),
	)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	answerLikeController.Like(
		recorder,
		request.WithContext(ctx),
		httprouter.Params{},
	)

	webResp := response.WebResponse{}
	dec := json.NewDecoder(recorder.Result().Body)
	dec.Decode(&webResp)

	assert.Equal(t, http.StatusOK, webResp.Code)

	data := webResp.Data.(map[string]interface{})
	assert.Equal(
		t,
		dislikedAnswer[0].ID,
		uint(data["answerId"].(float64)),
	)
	assert.Equal(
		t,
		false,
		data["isLiked"].(bool),
	)

	mockLike.Unset()
	mockGetAnswer.Unset()

}

func TestUnauthorized(t *testing.T) {

	bodyBytes, _ := json.Marshal(&likeRequestBody)
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
			answerLikeController.Like(
				recorder,
				request,
				httprouter.Params{},
			)
		},
	)

}
