package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type AnswerLikeRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, answerLike entity.AnswerLike) (entity.AnswerLike, error)
	FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerLikeParam) ([]entity.AnswerLike, error)
	Delete(ctx context.Context, DB *gorm.DB, answerLike entity.AnswerLike) (entity.AnswerLike, error)
}
