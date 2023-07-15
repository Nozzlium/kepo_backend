package questionlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/exception"

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

func (service *QuestionLikeServiceImpl) AssignLike(ctx context.Context, params param.QuestionLikeParam) (response.QuestionLikeResponse, error) {
	var err error
	var resp response.QuestionLikeResponse
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

	res, err := service.QuestionRepository.FindDetailed(
		ctx,
		service.DB,
		param.QuestionParam{
			UserID: params.QuestionLike.UserID,
			Question: entity.Question{
				ID: params.QuestionLike.QuestionID,
			},
		},
	)
	if len(res) < 1 {
		return resp, exception.NotFoundError{}
	}
	likedQuestion := res[0]
	resp = response.QuestionLikeResponse{
		QuestionID: likedQuestion.ID,
		Likes:      likedQuestion.Likes,
		IsLiked:    likedQuestion.UserLiked != 0,
	}
	return resp, err
}
