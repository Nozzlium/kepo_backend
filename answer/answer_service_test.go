package answer

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/repositorymock"
	"nozzlium/kepo_backend/data/repository/result"
	"nozzlium/kepo_backend/data/response"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var answerRepositoryMock = repositorymock.AnswerRepositoryMock{Mock: mock.Mock{}}
var answerService = AnswerServiceImpl{
	AnswerRepository: &answerRepositoryMock,
}

var expectedAnswers = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      10,
		IsLiked:    0,
	},
	{
		ID:         2,
		QuestionID: 1,
		UserID:     2,
		Username:   "user2",
		Likes:      5,
		IsLiked:    1,
	},
}

func TestServiceCreateAnswer(t *testing.T) {
	mockCall := answerRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.Answer{
			ID:         1,
			UserID:     1,
			Content:    "test",
			QuestionID: 1,
		},
		nil,
	)
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
	mockCall := answerRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(entity.Answer{}, gorm.ErrInvalidField)
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
	mockCall := answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedAnswers,
		nil,
	)
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
