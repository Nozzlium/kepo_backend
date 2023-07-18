package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/result"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, question entity.Question) (entity.Question, error)
	Find(ctx context.Context, DB *gorm.DB, param param.QuestionParam) ([]entity.Question, error)
	FindOneBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (entity.Question, error)
	FindDetailed(ctx context.Context, DB *gorm.DB, param param.QuestionParam) ([]result.QuestionResult, error)
	FindOneDetailedBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (result.QuestionResult, error)
}
