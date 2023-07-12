package question

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/result"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var questionService QuestionService = &QuestionServiceImpl{QuestionRepository: &questionRepositoryMock}

func TestCreateQuestionSuccess(t *testing.T) {
	mockCall := questionRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedInsertedQuestion,
		nil,
	)
	mockCallQuestion := mockReturnOneDetailed()
	question, err := questionService.CreateQuestion(
		context.Background(),
		entity.Question{
			UserID:      1,
			CategoryID:  1,
			Content:     "test",
			Description: "test",
		},
	)
	mockCall.Unset()
	mockCallQuestion.Unset()

	assert.Nil(t, err)
	assert.Equal(t, uint(expectedQuestions[0].ID), question.ID)
	assert.Equal(t, uint(expectedQuestions[0].UserID), question.User.ID)
	assert.Equal(t, uint(expectedQuestions[0].CategoryID), question.Category.ID)
}

func TestCreateQuestionFail(t *testing.T) {
	mockCall := questionRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.Question{},
		gorm.ErrInvalidField,
	)
	_, err := questionService.CreateQuestion(
		context.Background(),
		entity.Question{
			UserID:      1,
			CategoryID:  1,
			Content:     "test",
			Description: "test",
		},
	)
	mockCall.Unset()
	assert.NotNil(t, err)
}

func TestGetQuestionsSuccess(t *testing.T) {
	mockCall := questionRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedQuestions,
		nil,
	)
	resp, err := questionService.FindAll(
		context.Background(),
		param.QuestionParam{},
	)
	mockCall.Unset()
	assert.Nil(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, uint(1), resp[0].ID)
	assert.Equal(t, uint(2), resp[1].User.ID)
}

func TestGetQuestionsError(t *testing.T) {
	mockCall := questionRepositoryMock.Mock.On(
		"FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.QuestionResult{},
		gorm.ErrRecordNotFound,
	)
	_, err := questionService.FindAll(
		context.Background(),
		param.QuestionParam{},
	)
	mockCall.Unset()

	assert.NotNil(t, err)
}

func TestGetOneSuccess(t *testing.T) {
	mockCall := questionRepositoryMock.Mock.On(
		"FindOneDetailedBy",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedQuestions[0],
		nil,
	)
	res, err := questionService.FindOneBy(
		context.Background(),
		param.QuestionParam{
			Question: entity.Question{
				ID: 1,
			},
		},
	)
	mockCall.Unset()

	assert.Nil(t, err)
	assert.Equal(t, uint(1), res.ID)
}

func TestGetOneError(t *testing.T) {
	mockCall := questionRepositoryMock.Mock.On(
		"FindOneDetailedBy",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		result.QuestionResult{},
		gorm.ErrRecordNotFound,
	)
	_, err := questionService.FindOneBy(
		context.Background(),
		param.QuestionParam{
			Question: entity.Question{
				ID: 1,
			},
		},
	)
	mockCall.Unset()

	assert.NotNil(t, err)
}
