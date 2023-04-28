package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/result"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error)
	FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]entity.Answer, error)
	FindOneBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (entity.Answer, error)
	FindDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]result.AnswerResult, error)
}
