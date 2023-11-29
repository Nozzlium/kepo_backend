package questionlike

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

type QuestionLikeServiceImpl struct {
	QuestionLikeRepository repository.QuestionLikeRepository
	QuestionRepository     repository.QuestionRepository
	UserRepository         repository.UserRepository
	NotificationRepository repository.NotificationRepository
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

	var resQuestion result.QuestionResult
	var resUser entity.User

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		resQuestion, err = service.QuestionRepository.FindOneDetailedBy(
			ctx,
			service.DB,
			param.QuestionParam{
				UserID: params.QuestionLike.UserID,
				Question: entity.Question{
					ID: params.QuestionLike.QuestionID,
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
				ID: params.QuestionLike.UserID,
			},
		)
		wg.Done()
	}()

	wg.Wait()

	if err == nil {
		service.NotificationRepository.Create(
			ctx,
			service.DB,
			entity.Notification{
				UserID:     resQuestion.UserID,
				QuestionID: resQuestion.ID,
				NotifType:  constants.NOTIFICATION_TYPE_LIKE,
				Headline:   fmt.Sprintf(constants.QUESTION_LIKE_NOTIF, resUser.Username),
				Preview:    helper.GetNotificationPreview(resQuestion.Content),
			},
		)
	} else {
		return response.QuestionResponse{}, err
	}

	return helper.QuestionResultToResponse(resQuestion), err
}
