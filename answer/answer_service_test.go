package answer

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceCreateAnswer(t *testing.T) {
	mockCall := mockReturnCreatedAnswer()
	response, err := answerService.CreateAnswer(
		context.Background(),
		entity.Answer{
			UserID:     1,
			Content:    "test",
			QuestionID: 1,
		},
	)
	mockCall.Unset()
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, uint(1), response.User.ID)
	assert.Nil(t, err)
}

func TestFailCreateAnswer(t *testing.T) {
	mockCall := mockCreateAnswerError()
	_, err := answerService.CreateAnswer(
		context.Background(),
		entity.Answer{
			QuestionID: 1,
			UserID:     1,
			Content:    "test",
		},
	)
	mockCall.Unset()
	assert.NotNil(t, err)
}

func TestGetAnswers(t *testing.T) {
	mockCall := mockReturnAnswers()
	resp, err := answerService.FindBy(
		context.Background(),
		param.AnswerParam{},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	for _, ansResp := range resp {
		assert.IsType(t, response.AnswerResponse{}, ansResp)
	}
	assert.Equal(t, uint(1), resp[0].ID)
	assert.Equal(t, uint(2), resp[1].ID)
}

func TestGetAnswersError(t *testing.T) {
	mockCall := mockReturnAnswersError()

	_, err := answerService.FindBy(
		context.Background(),
		param.AnswerParam{},
	)
	assert.NotNil(t, err)

	mockCall.Unset()
}
