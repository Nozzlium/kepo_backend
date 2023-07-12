package question

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository/repositorymock"
	"nozzlium/kepo_backend/data/repository/result"
	"nozzlium/kepo_backend/data/requestbody"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var questionRepositoryMock = repositorymock.QuestionRepositoryMock{Mock: &mock.Mock{}}

var expectedQuestions = []result.QuestionResult{
	{
		ID:         1,
		UserID:     1,
		Username:   "User1",
		CategoryID: 1,
	},
	{
		ID:         2,
		UserID:     2,
		Username:   "User2",
		CategoryID: 2,
	},
}

var expectedQuestionsSameUser = []result.QuestionResult{
	{
		ID:       1,
		UserID:   1,
		Username: "User1",
	},
	{
		ID:       3,
		UserID:   1,
		Username: "User1",
	},
}

var expectedInsertedQuestion = entity.Question{
	ID:          1,
	UserID:      1,
	CategoryID:  1,
	Content:     "test",
	Description: "test",
}

var createRequestBody = requestbody.Question{
	CategoryID:  1,
	Content:     "test",
	Description: "test",
}

func mockReturnInsertSuccess() *mock.Call {
	return questionRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedInsertedQuestion,
		nil,
	)
}

func mockReturnInsertError() *mock.Call {
	return questionRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.Question{},
		gorm.ErrInvalidField,
	)
}

func mockReturnQuestionsSuccess() *mock.Call {
	return questionRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedQuestions,
		nil,
	)
}

func mockReturnQuestionsSameUserSuccess() *mock.Call {
	return questionRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedQuestionsSameUser,
		nil,
	)
}

func mockReturnQuestionError() *mock.Call {
	return questionRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.QuestionResult{},
		gorm.ErrRecordNotFound,
	)
}

func mockReturnOneQuestionSuccess() *mock.Call {
	return questionRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedQuestions[0:1],
		nil,
	)
}

func mockReturnEmptyQuestion() *mock.Call {
	return questionRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.QuestionResult{},
		nil,
	)
}

func mockReturnOneDetailed() *mock.Call {
	return questionRepositoryMock.Mock.On(
		"FindOneDetailedBy",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedQuestions[0],
		nil,
	)
}

func mockReturnOneNotFound() *mock.Call {
	return questionRepositoryMock.Mock.On(
		"FindOneDetailedBy",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		result.QuestionResult{},
		gorm.ErrRecordNotFound,
	)
}
