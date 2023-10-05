package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type AnswerLikeServiceImpl struct {
	AnswerLikeRepository repository.AnswerLikeRepository
	AnswerRepository     repository.AnswerRepository
	DB                   *gorm.DB
}

func NewAnswerLikeService(
	answerLikeRepository repository.AnswerLikeRepository,
	answerRepository repository.AnswerRepository,
	DB *gorm.DB,
) *AnswerLikeServiceImpl {
	return &AnswerLikeServiceImpl{
		AnswerLikeRepository: answerLikeRepository,
		AnswerRepository:     answerRepository,
		DB:                   DB,
	}
}

func (service *AnswerLikeServiceImpl) AssignLike(ctx context.Context, params param.AnswerLikeParam) (response.AnswerResponse, error) {
	var err error
	if params.IsLike {
		_, err = service.AnswerLikeRepository.Insert(
			ctx,
			service.DB,
			params.AnswerLike,
		)
	} else {
		_, err = service.AnswerLikeRepository.Delete(
			ctx,
			service.DB,
			params.AnswerLike,
		)
	}
	if err != nil {
		return response.AnswerResponse{}, err
	}
	res, err := service.AnswerRepository.FindOneDetailed(
		ctx,
		service.DB,
		param.AnswerParam{
			PaginationParam: param.InitPaginationParam(),
			UserID:          params.AnswerLike.UserID,
			Answer: entity.Answer{
				ID: params.AnswerLike.AnswerID,
			},
		},
	)
	if err != nil {
		return response.AnswerResponse{}, err
	}
	likedAnswer := res
	return helper.AnswerResultToResponse(likedAnswer), nil
}
