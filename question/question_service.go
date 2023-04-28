package question

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type QuestionService interface {
	CreateQuestion(ctx context.Context, question entity.Question) (entity.Question, error)
	FindAll(ctx context.Context, param param.QuestionParam) ([]response.QuestionResponse, error)
	FindOneBy(ctx context.Context, param param.QuestionParam) (response.QuestionResponse, error)
}
