package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type QuestionLikeRepository interface {
	Insert(ctx context.Context, DB *gorm.DB, questionLike entity.QuestionLike) (entity.QuestionLike, error)
	FindBy(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) ([]entity.QuestionLike, error)
	FindOneBy(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) (entity.QuestionLike, error)
	Delete(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) (entity.QuestionLike, error)
}
