package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"

	"gorm.io/gorm"
)

type AnswerLikeServiceImpl struct {
	AnswerLikeRepository repository.AnswerLikeRepository
	AnswerRepository     repository.AnswerRepository
	DB                   *gorm.DB
}

func (service *AnswerLikeServiceImpl) AssignLike(ctx context.Context, params param.AnswerLikeParam) (response.AnswerLikeResponse, error) {
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
		return response.AnswerLikeResponse{}, err
	}
	res, err := service.AnswerRepository.FindDetailed(
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
		return response.AnswerLikeResponse{}, err
	}
	if len(res) < 1 {
		return response.AnswerLikeResponse{}, exception.NotFoundError{}
	}
	likedAnswer := res[0]
	return response.AnswerLikeResponse{
		AnswerID: likedAnswer.ID,
		Likes:    likedAnswer.Likes,
		IsLiked:  likedAnswer.IsLiked != 0,
	}, nil
}
