package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type AnswerLikeServiceImpl struct {
	AnswerLikeRepository repository.AnswerLikeRepository
	DB                   *gorm.DB
}

func (service *AnswerLikeServiceImpl) AssignLike(ctx context.Context, param param.AnswerLikeParam) (response.AnswerLikeResponse, error) {
	if param.IsLike {
		_, err := service.AnswerLikeRepository.Insert(
			ctx,
			service.DB,
			param.AnswerLike,
		)
		return helper.AnswerLikeParamToResponse(param), err
	} else {
		_, err := service.AnswerLikeRepository.Delete(
			ctx,
			service.DB,
			param.AnswerLike,
		)
		return helper.AnswerLikeParamToResponse(param), err
	}
}
