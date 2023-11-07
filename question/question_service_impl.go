package question

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/mysqlerr"

	"gorm.io/gorm"
)

type QuestionServiceImpl struct {
	QuestionRepository repository.QuestionRepository
	DB                 *gorm.DB
}

func NewQuestionService(
	questionRepository repository.QuestionRepository,
	DB *gorm.DB,
) *QuestionServiceImpl {
	return &QuestionServiceImpl{
		QuestionRepository: questionRepository,
		DB:                 DB,
	}
}

func (service *QuestionServiceImpl) CreateQuestion(ctx context.Context, question entity.Question) (response.QuestionResponse, error) {
	inserted, err := service.QuestionRepository.Insert(
		ctx,
		service.DB,
		question,
	)
	if err != nil {
		return response.QuestionResponse{}, err
	}
	newQuestion, err := service.QuestionRepository.FindOneDetailedBy(
		ctx,
		service.DB,
		param.QuestionParam{
			Question: entity.Question{
				ID: inserted.ID,
			},
		},
	)
	return helper.QuestionResultToResponse(newQuestion), err
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
	err = mysqlerr.CheckMySQLError(err)
	return helper.QuestionResultToResponse(question), err
}

func (service *QuestionServiceImpl) FindLikedByUser(ctx context.Context, param param.LikedQuestionParam) ([]response.QuestionResponse, error) {
	questions, err := service.QuestionRepository.FindDetailedLikedByUser(
		ctx,
		service.DB,
		param,
	)
	return helper.QuestionResultsToResponses(questions), err
}

func (service *QuestionServiceImpl) Delete(ctx context.Context, question entity.Question) (response.QuestionResponse, error) {
	toBeDeleted, err := service.QuestionRepository.FindOneDetailedBy(
		ctx,
		service.DB,
		param.QuestionParam{
			Question: entity.Question{
				ID: question.ID,
			},
		},
	)
	if err != nil {
		return response.QuestionResponse{}, err
	}
	_, err = service.QuestionRepository.Delete(ctx, service.DB, question)
	return helper.QuestionResultToResponse(toBeDeleted), err
}

func (service *QuestionServiceImpl) Update(ctx context.Context, question entity.Question) (response.QuestionResponse, error) {
	update, err := service.QuestionRepository.Update(ctx, service.DB, question)
	if err != nil {
		return response.QuestionResponse{}, err
	}
	result, err := service.QuestionRepository.FindOneDetailedBy(ctx, service.DB, param.QuestionParam{Question: entity.Question{ID: update.ID}})
	return helper.QuestionResultToResponse(result), err
}
