package questionlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type QuestionLikeServiceImpl struct {
	QuestionLikeRepository repository.QuestionLikeRepository
	QuestionRepository     repository.QuestionRepository
	DB                     *gorm.DB
}

func NewQuestionLikeService(
	questionLikeRepository repository.QuestionLikeRepository,
	questionRepository repository.QuestionRepository,
	DB *gorm.DB,
) *QuestionLikeServiceImpl {
	return &QuestionLikeServiceImpl{
		QuestionLikeRepository: questionLikeRepository,
		QuestionRepository:     questionRepository,
		DB:                     DB,
	}
}

func (service *QuestionLikeServiceImpl) AssignLike(ctx context.Context, params param.QuestionLikeParam) (response.QuestionResponse, error) {
	var err error
	var resp response.QuestionResponse
	if params.IsLiked {
		_, err = service.QuestionLikeRepository.Insert(
			ctx,
			service.DB,
			params.QuestionLike,
		)
	} else {
		_, err = service.QuestionLikeRepository.Delete(
			ctx,
			service.DB,
			params,
		)
	}
	if err != nil {
		return resp, err
	}

	res, err := service.QuestionRepository.FindOneDetailedBy(
		ctx,
		service.DB,
		param.QuestionParam{
			UserID: params.QuestionLike.UserID,
			Question: entity.Question{
				ID: params.QuestionLike.QuestionID,
			},
		},
	)
	return helper.QuestionResultToResponse(res), err
}
