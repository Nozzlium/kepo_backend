package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository"

	"gorm.io/gorm"
)

type AnswerLikeServiceImpl struct {
	AnswerLikeRepository repository.AnswerLikeRepository
	DB                   *gorm.DB
}

func (service *AnswerLikeServiceImpl) AssignLike(ctx context.Context, param param.AnswerLikeParam) (entity.AnswerLike, error) {
	panic("not implemented") // TODO: Implement
}
