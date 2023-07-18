package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type AnswerLikeRepositoryMock struct {
	Mock *mock.Mock
}

func (repository *AnswerLikeRepositoryMock) Insert(ctx context.Context, DB *gorm.DB, answerLike entity.AnswerLike) (entity.AnswerLike, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		answerLike,
	)
	args0 := args[0].(entity.AnswerLike)
	err := args[1]
	if err == nil {
		return args0, nil
	}
	return args0, err.(error)
}

func (repository *AnswerLikeRepositoryMock) FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerLikeParam) ([]entity.AnswerLike, error) {
	return []entity.AnswerLike{}, nil
}

func (repository *AnswerLikeRepositoryMock) Delete(ctx context.Context, DB *gorm.DB, answerLike entity.AnswerLike) (entity.AnswerLike, error) {
	args := repository.Mock.Called(
		ctx,
		DB,
		answerLike,
	)
	args0 := args[0].(entity.AnswerLike)
	err := args[1]
	if err == nil {
		return args0, nil
	}
	return args0, err.(error)
}
