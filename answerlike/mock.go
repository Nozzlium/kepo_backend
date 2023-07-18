package answerlike

import (
	"errors"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/result"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
)

var answerLikeRepositoryMock = repository.AnswerLikeRepositoryMock{Mock: &mock.Mock{}}
var answerRepositoryMock = repository.AnswerRepositoryMock{Mock: &mock.Mock{}}
var answerLikeService = AnswerLikeServiceImpl{
	AnswerLikeRepository: &answerLikeRepositoryMock,
	AnswerRepository:     &answerRepositoryMock,
}
var answerLikeController AnswerLikeController = &AnswerLikeControllerImpl{
	AnswerLikeService: &answerLikeService,
	Validator:         validator.New(),
}

var likeRequestBody = requestbody.AnswerLike{
	AnswerID: 1,
	IsLiked:  true,
}
var dislikeRequestBody = requestbody.AnswerLike{
	AnswerID: 1,
	IsLiked:  false,
}

var likedAnswer = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Likes:      1,
		IsLiked:    1,
	},
}

var dislikedAnswer = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Likes:      0,
		IsLiked:    0,
	},
}

func mockInsertLike() *mock.Call {
	return answerLikeRepositoryMock.Mock.On(
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
}

func mockRemoveLike() *mock.Call {
	return answerLikeRepositoryMock.Mock.On(
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
}

func mockGetLikedAnswer() *mock.Call {
	return answerRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		likedAnswer,
		nil,
	)
}

func mockGetDislikedAnswer() *mock.Call {
	return answerRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		dislikedAnswer,
		nil,
	)
}

func mockLikeRepositoryError() *mock.Call {
	return answerLikeRepositoryMock.Mock.On(
		"Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.AnswerLike{},
		errors.New("error"),
	)
}

func mockGetLikedAnswerError() *mock.Call {
	return answerRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.AnswerResult{},
		errors.New("error"),
	)
}

func mockGetLikedAnswerEmpty() *mock.Call {
	return answerRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.AnswerResult{},
		nil,
	)
}
