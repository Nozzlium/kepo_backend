package question

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type QuestionServiceImpl struct {
	QuestionRepository repository.QuestionRepository
	DB                 *gorm.DB
}

func (service *QuestionServiceImpl) CreateQuestion(ctx context.Context, question entity.Question) (entity.Question, error) {
	inserted, err := service.QuestionRepository.Insert(
		ctx,
		service.DB,
		question,
	)
	return inserted, err
}

func (service *QuestionServiceImpl) FindAll(ctx context.Context, param param.QuestionParam) ([]response.QuestionResponse, error) {
	questions, err := service.QuestionRepository.FindDetailed(
		ctx,
		service.DB,
		param,
	)
	return helper.QuestionResultsToResponses(questions), err
}

func (service *QuestionServiceImpl) FindOneBy(ctx context.Context, param param.QuestionParam) (response.QuestionResponse, error) {
	question, err := service.QuestionRepository.FindOneDetailedBy(
		ctx,
		service.DB,
		param,
	)
	return helper.QuestionResultToResponse(question), err
}
