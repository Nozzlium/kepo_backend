package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
)

type AnswerLikeService interface {
	AssignLike(ctx context.Context, param param.AnswerLikeParam) (entity.AnswerLike, error)
}
