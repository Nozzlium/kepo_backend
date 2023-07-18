package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type QuestionLikeRepositoryMock struct {
	Mock *mock.Mock
}

func (repository *QuestionLikeRepositoryMock) Insert(ctx context.Context, DB *gorm.DB, questionLike entity.QuestionLike) (entity.QuestionLike, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		questionLike,
	)
	args0 := args[0].(entity.QuestionLike)
	args1 := args[1]
	if args1 != nil {
		return args0, args1.(error)
	}
	return args0, nil
}

func (repository *QuestionLikeRepositoryMock) FindBy(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) ([]entity.QuestionLike, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *QuestionLikeRepositoryMock) FindOneBy(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) (entity.QuestionLike, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *QuestionLikeRepositoryMock) Delete(ctx context.Context, DB *gorm.DB, param param.QuestionLikeParam) (entity.QuestionLike, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		param,
	)
	args0 := args[0].(entity.QuestionLike)
	args1 := args[1]
	if args1 != nil {
		return args0, args1.(error)
	}
	return args0, nil
}
