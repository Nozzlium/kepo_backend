package answer

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type AnswerService interface {
	CreateAnswer(ctx context.Context, answer entity.Answer) (response.AnswerResponse, error)
	FindBy(ctx context.Context, param param.AnswerParam) ([]response.AnswerResponse, error)
}
