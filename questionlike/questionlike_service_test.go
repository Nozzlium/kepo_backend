package questionlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserLike(t *testing.T) {
	mockCall := mockLikeQuestion()
	mockCallLikedQuestion := mockGetLikedAnswer()
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
	mockCallLikedQuestion.Unset()

	assert.Nil(t, err)
	assert.Equal(t, true, response.IsLiked)
	assert.Equal(t, uint(1), response.ID)
}

func TestAssignDislike(t *testing.T) {
	mockCall := mockDislikeQuestion()
	mockCallDislikedQuestion := mockGetDislikedAnswer()
	response, err := questionLikeService.AssignLike(
		context.Background(),
		param.QuestionLikeParam{
			IsLiked: false,
			QuestionLike: entity.QuestionLike{
				UserID:     1,
				QuestionID: 2,
			},
		},
	)
	mockCall.Unset()
	mockCallDislikedQuestion.Unset()

	assert.Nil(t, err)
	assert.Equal(t, uint(2), response.ID)
}
