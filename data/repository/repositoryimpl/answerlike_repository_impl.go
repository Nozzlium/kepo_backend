package repositoryimpl

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type AnswerLikeRepositoryImpl struct {
}

func (repository *AnswerLikeRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, answerLike entity.AnswerLike) (entity.AnswerLike, error) {
	create := DB.WithContext(ctx).Create(&answerLike)
	return answerLike, create.Error
}

func (repository *AnswerLikeRepositoryImpl) FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerLikeParam) ([]entity.AnswerLike, error) {
	answerLikes := []entity.AnswerLike{}
	find := DB.WithContext(ctx).Where(&param.AnswerLike).Find(&answerLikes)
	return answerLikes, find.Error
}

func (repository *AnswerLikeRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, answerLike entity.AnswerLike) (entity.AnswerLike, error) {
	delete := DB.WithContext(ctx).Delete(answerLike)
	return answerLike, delete.Error
}
