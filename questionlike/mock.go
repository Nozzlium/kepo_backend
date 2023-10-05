package questionlike

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/result"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
)

var questionLikeRepositoryMock = repository.QuestionLikeRepositoryMock{Mock: &mock.Mock{}}
var questionRepositoryMock = repository.QuestionRepositoryMock{Mock: &mock.Mock{}}
var questionLikeService = QuestionLikeServiceImpl{
	QuestionLikeRepository: &questionLikeRepositoryMock,
	QuestionRepository:     &questionRepositoryMock,
}
var questionLikeController QuestionLikeController = &QuestionLikeControllerImpl{
	QuestionLikeService: &questionLikeService,
	Validator:           validator.New(),
}

var likeRequestBody = requestbody.QuestionLike{
	QuestionID: 1,
	IsLiked:    true,
}

var dislikeRequestBody = requestbody.QuestionLike{
	QuestionID: 2,
	IsLiked:    false,
}

func mockLikeQuestion() *mock.Call {
	return questionLikeRepositoryMock.Mock.On(
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
}

func mockDislikeQuestion() *mock.Call {
	return questionLikeRepositoryMock.Mock.On(
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
}

var expectedLikedQuestion = []result.QuestionResult{
	{
		ID:        1,
		UserID:    1,
		Username:  "User1",
		Likes:     25,
		UserLiked: 1,
	},
}

var expectedDislikedQuestion = []result.QuestionResult{
	{
		ID:        2,
		UserID:    1,
		Username:  "User1",
		Likes:     12,
		UserLiked: 0,
	},
}

func mockGetLikedAnswer() *mock.Call {
	return questionRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedLikedQuestion,
		nil,
	)
}

func mockGetDislikedAnswer() *mock.Call {
	return questionRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedDislikedQuestion,
		nil,
	)
}
