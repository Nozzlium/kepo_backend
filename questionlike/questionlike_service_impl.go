package questionlike

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type QuestionLikeServiceImpl struct {
	QuestionLikeRepository repository.QuestionLikeRepository
	DB                     *gorm.DB
}

func (service *QuestionLikeServiceImpl) AssignLike(ctx context.Context, param param.QuestionLikeParam) (response.QuestionLikeResponse, error) {
	if param.IsLiked {
		_, err := service.QuestionLikeRepository.Insert(
			ctx,
			service.DB,
			param.QuestionLike,
		)
		return helper.QuestionLikeParamToResponse(param), err
	} else {
		_, err := service.QuestionLikeRepository.Delete(
			ctx,
			service.DB,
			param,
		)
		return helper.QuestionLikeParamToResponse(param), err
	}
}
