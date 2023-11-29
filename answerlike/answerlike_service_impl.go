package answerlike

import (
	"context"
	"fmt"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/data/result"
	"nozzlium/kepo_backend/helper"
	"sync"

	"gorm.io/gorm"
)

type AnswerLikeServiceImpl struct {
	AnswerLikeRepository   repository.AnswerLikeRepository
	AnswerRepository       repository.AnswerRepository
	UserRepository         repository.UserRepository
	NotificationRepository repository.NotificationRepository
	DB                     *gorm.DB
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

	var resAnswer result.AnswerResult
	var resUser entity.User

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		resAnswer, err = service.AnswerRepository.FindOneDetailed(
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
		wg.Done()
	}()

	go func() {
		resUser, err = service.UserRepository.FindOneBasedOnIdentity(
			ctx,
			service.DB,
			entity.User{
				ID: params.AnswerLike.UserID,
			},
		)
		wg.Done()
	}()

	if err == nil {
		service.NotificationRepository.Create(
			ctx,
			service.DB,
			entity.Notification{
				UserID:     resAnswer.UserID,
				QuestionID: resAnswer.QuestionID,
				NotifType:  constants.NOTIFICATION_TYPE_LIKE,
				Headline:   fmt.Sprintf(constants.ANSWER_LIKE_NOTIF, resUser.Username),
				Preview:    helper.GetNotificationPreview(resAnswer.Content),
			},
		)
	} else {
		return response.AnswerResponse{}, err
	}

	return helper.AnswerResultToResponse(resAnswer), nil
}
