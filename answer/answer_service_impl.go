package answer

import (
	"context"
	"fmt"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type AnswerServiceImpl struct {
	AnswerRepository       repository.AnswerRepository
	NotificationRepository repository.NotificationRepository
	DB                     *gorm.DB
}

func NewAnswerService(
	answerRepository repository.AnswerRepository,
	notificationRepository repository.NotificationRepository,
	DB *gorm.DB,
) *AnswerServiceImpl {
	return &AnswerServiceImpl{
		AnswerRepository:       answerRepository,
		NotificationRepository: notificationRepository,
		DB:                     DB,
	}
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
	if res.UserID != res.QuestionPosterID {
		newNotif := entity.Notification{
			UserID:     res.QuestionPosterID,
			QuestionID: res.QuestionID,
			NotifType:  constants.NOTIFICATION_TYPE_ANSWER,
			Headline:   fmt.Sprintf(constants.ANSWER_POSTED_NOTIF, res.Username),
			Preview:    helper.GetNotificationPreview(res.Content),
		}
		service.NotificationRepository.Create(
			ctx,
			service.DB,
			newNotif,
		)
	}
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

func (service *AnswerServiceImpl) Delete(ctx context.Context, answer entity.Answer) (response.AnswerResponse, error) {
	toBeDeleted, err := service.AnswerRepository.FindOneDetailed(
		ctx,
		service.DB,
		param.AnswerParam{
			Answer: entity.Answer{
				ID: answer.ID,
			},
		},
	)
	if err != nil {
		return response.AnswerResponse{}, err
	}
	_, err = service.AnswerRepository.Delete(ctx, service.DB, answer)
	return helper.AnswerResultToResponse(toBeDeleted), err
}

func (service *AnswerServiceImpl) Update(ctx context.Context, answer entity.Answer) (response.AnswerResponse, error) {
	_, err := service.AnswerRepository.Update(ctx, service.DB, answer)
	if err != nil {
		return response.AnswerResponse{}, err
	}
	result, err := service.AnswerRepository.FindOneDetailed(
		ctx,
		service.DB,
		param.AnswerParam{
			UserID: answer.UserID,
			Answer: entity.Answer{
				ID: answer.ID,
			},
		},
	)
	return helper.AnswerResultToResponse(result), err
}
