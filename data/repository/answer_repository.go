package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/result"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error)
	FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]entity.Answer, error)
	FindOneBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (entity.Answer, error)
	FindDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]result.AnswerResult, error)
	FindOneDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (result.AnswerResult, error)
	Delete(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error)
	Update(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error)
}
