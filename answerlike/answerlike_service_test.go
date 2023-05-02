package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/repositorymock"
	"nozzlium/kepo_backend/data/response"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var answerLikeRepositoryMock = repositorymock.AnswerLikeRepositoryMock{Mock: &mock.Mock{}}
var answerLikeService = AnswerLikeServiceImpl{
	AnswerLikeRepository: &answerLikeRepositoryMock,
}

func TestAssignLike(t *testing.T) {
	mockCall := answerLikeRepositoryMock.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.AnswerLike{
			UserID:   1,
			AnswerID: 1,
		},
		nil,
	)
	answer, err := answerLikeService.AssignLike(
		context.Background(),
		param.AnswerLikeParam{
			IsLike: true,
			AnswerLike: entity.AnswerLike{
				UserID:   1,
				AnswerID: 1,
			},
		},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.IsType(t, response.AnswerLikeResponse{}, answer)
	assert.Equal(t, true, answer.IsLike)
	assert.Equal(t, uint(1), answer.UserID)
	assert.Equal(t, uint(1), answer.AnswerID)
}

func TestAssignDislike(t *testing.T) {
	mockCall := answerLikeRepositoryMock.Mock.On(
		"Delete",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.AnswerLike{
			UserID:   1,
			AnswerID: 1,
		},
		nil,
	)
	answer, err := answerLikeService.AssignLike(
		context.Background(),
		param.AnswerLikeParam{
			IsLike: false,
			AnswerLike: entity.AnswerLike{
				UserID:   1,
				AnswerID: 1,
			},
		},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.IsType(t, response.AnswerLikeResponse{}, answer)
	assert.Equal(t, false, answer.IsLike)
	assert.Equal(t, uint(1), answer.UserID)
	assert.Equal(t, uint(1), answer.AnswerID)
}
