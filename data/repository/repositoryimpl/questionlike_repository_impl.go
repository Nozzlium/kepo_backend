package repositoryimpl

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"gorm.io/gorm"
)

type QuestionLikeRepositoryImpl struct {
}

func (repository *QuestionLikeRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, questionLike entity.QuestionLike) (entity.QuestionLike, error) {
	insert := DB.WithContext(ctx).Create(&questionLike)
	return questionLike, insert.Error
}

func (repository *QuestionLikeRepositoryImpl) FindBy(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) ([]entity.QuestionLike, error) {
	questionLikes := []entity.QuestionLike{}
	find := DB.WithContext(ctx).Where(param.QuestionLike).Find(&questionLikes)
	return questionLikes, find.Error
}

func (repository *QuestionLikeRepositoryImpl) FindOneBy(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) (entity.QuestionLike, error) {
	questionLike := entity.QuestionLike{}
	find := DB.WithContext(ctx).Where(param.QuestionLike).First(&questionLike)
	return questionLike, find.Error
}

func (repository *QuestionLikeRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) (entity.QuestionLike, error) {
	delete := DB.WithContext(ctx).Delete(param.QuestionLike)
	return param.QuestionLike, delete.Error
}
