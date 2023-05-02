package questionlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/repositorymock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var questionLikeRepositoryMock = repositorymock.QuestionLikeRepositoryMock{Mock: &mock.Mock{}}
var questionLikeService = QuestionLikeServiceImpl{
	QuestionLikeRepository: &questionLikeRepositoryMock,
}

func TestAddUserLike(t *testing.T) {
	mockCall := questionLikeRepositoryMock.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.QuestionLike{
			UserID:     1,
			QuestionID: 1,
		},
		nil,
	)
	response, err := questionLikeService.AssignLike(
		context.Background(),
		param.QuestionLikeParam{
			IsLiked: true,
			QuestionLike: entity.QuestionLike{
				UserID:     1,
				QuestionID: 1,
			},
		},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.Equal(t, true, response.IsLiked)
	assert.Equal(t, uint(1), response.UserID)
	assert.Equal(t, uint(1), response.QuestionID)
}

func TestAssignDislike(t *testing.T) {
	mockCall := questionLikeRepositoryMock.Mock.On(
		"Delete",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.QuestionLike{
			UserID:     1,
			QuestionID: 1,
		},
		nil,
	)
	response, err := questionLikeService.AssignLike(
		context.Background(),
		param.QuestionLikeParam{
			IsLiked: false,
			QuestionLike: entity.QuestionLike{
				UserID:     1,
				QuestionID: 1,
			},
		},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.Equal(t, uint(1), response.UserID)
	assert.Equal(t, uint(1), response.QuestionID)
}
