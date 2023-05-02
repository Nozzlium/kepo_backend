package answerlike

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type AnswerLikeService interface {
	AssignLike(ctx context.Context, param param.AnswerLikeParam) (response.AnswerLikeResponse, error)
}
