package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignLike(t *testing.T) {
	mockCall := mockInsertLike()
	mockGetLikedAnswer := mockGetLikedAnswer()
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
	mockGetLikedAnswer.Unset()

	assert.Nil(t, err)
	assert.IsType(t, response.AnswerLikeResponse{}, answer)
	assert.Equal(t, true, answer.IsLiked)
	assert.Equal(t, uint(1), answer.ID)
}

func TestAssignDislike(t *testing.T) {
	mockCall := mockRemoveLike()
	mockDislikedAnswer := mockGetDislikedAnswer()
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
	mockDislikedAnswer.Unset()

	assert.Nil(t, err)
	assert.IsType(t, response.AnswerLikeResponse{}, answer)
	assert.Equal(t, false, answer.IsLiked)
	assert.Equal(t, uint(1), answer.ID)
}

func TestLikeServiceError(t *testing.T) {
	mockCall := mockLikeRepositoryError()
	mockDislikedAnswer := mockGetDislikedAnswer()
	_, err := answerLikeService.AssignLike(
		context.Background(),
		param.AnswerLikeParam{
			IsLike: true,
			AnswerLike: entity.AnswerLike{
				UserID:   1,
				AnswerID: 1,
			},
		},
	)
	assert.NotNil(t, err)
	mockCall.Unset()
	mockDislikedAnswer.Unset()
}

func TestAnswerRepoError(t *testing.T) {
	mockCall := mockInsertLike()
	mockDislikedAnswer := mockGetLikedAnswerError()
	_, err := answerLikeService.AssignLike(
		context.Background(),
		param.AnswerLikeParam{
			IsLike: true,
			AnswerLike: entity.AnswerLike{
				UserID:   1,
				AnswerID: 1,
			},
		},
	)
	assert.NotNil(t, err)
	mockCall.Unset()
	mockDislikedAnswer.Unset()
}

func TestGetLikedAnswerEmpty(t *testing.T) {
	mockCall := mockInsertLike()
	mockDislikedAnswer := mockGetLikedAnswerEmpty()
	_, err := answerLikeService.AssignLike(
		context.Background(),
		param.AnswerLikeParam{
			IsLike: true,
			AnswerLike: entity.AnswerLike{
				UserID:   1,
				AnswerID: 1,
			},
		},
	)
	assert.ErrorIs(
		t,
		err,
		exception.NotFoundError{},
	)
	mockCall.Unset()
	mockDislikedAnswer.Unset()
}
