package question

import (
	"context"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository/repositorymock"
	"nozzlium/kepo_backend/data/repository/result"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/tools"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var questionRepositoryMock = repositorymock.QuestionRepositoryMock{Mock: mock.Mock{}}

var expectedQuestions = []result.QuestionResult{
	{
		ID:       1,
		UserID:   1,
		Username: "User1",
	},
	{
		ID:       2,
		UserID:   2,
		Username: "User2",
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

var expectedOneQuestion = []result.QuestionResult{
	{
		ID:       1,
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

var claimContext = tools.JwtClaims{
	UserId: 1,
}

func getClaimContext(parent context.Context) context.Context {
	return context.WithValue(parent, constants.USER_ID_CLAIMS, claimContext)
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
		expectedOneQuestion,
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
