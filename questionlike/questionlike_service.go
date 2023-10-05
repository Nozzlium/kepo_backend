package questionlike

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type QuestionLikeService interface {
	AssignLike(ctx context.Context, param param.QuestionLikeParam) (response.QuestionResponse, error)
}
