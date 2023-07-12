package answer

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type AnswerServiceImpl struct {
	AnswerRepository repository.AnswerRepository
	DB               *gorm.DB
}

func (service *AnswerServiceImpl) CreateAnswer(ctx context.Context, answer entity.Answer) (response.AnswerResponse, error) {
	ans, err := service.AnswerRepository.Insert(
		ctx,
		service.DB,
		answer,
	)
	if err != nil {
		return response.AnswerResponse{}, err
	}
	res, err := service.AnswerRepository.FindOneDetailed(
		ctx,
		service.DB,
		param.AnswerParam{
			Answer: entity.Answer{
				ID: ans.ID,
			},
		},
	)
	return helper.AnswerResultToResponse(res), err
}

func (service *AnswerServiceImpl) FindBy(ctx context.Context, param param.AnswerParam) ([]response.AnswerResponse, error) {
	answers, err := service.AnswerRepository.FindDetailed(
		ctx,
		service.DB,
		param,
	)
	return helper.AnswersResultSliceToResponsesSlice(answers), err
}

func (service *AnswerServiceImpl) FindOneBy(ctx context.Context, param param.AnswerParam) (response.AnswerResponse, error) {
	answer, err := service.AnswerRepository.FindOneDetailed(
		ctx,
		service.DB,
		param,
	)
	return helper.AnswerResultToResponse(answer), err
}
