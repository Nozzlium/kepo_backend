package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/result"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type QuestionRepositoryMock struct {
	Mock *mock.Mock
}

func (repository *QuestionRepositoryMock) Insert(ctx context.Context, DB *gorm.DB, question entity.Question) (entity.Question, error) {
	args := repository.Mock.Called(
		ctx, DB, question,
	)
	args0 := args.Get(0)
	args1 := args.Get(1)
	if args1 == nil {
		return args0.(entity.Question), nil
	}
	return args0.(entity.Question), args1.(error)
}

func (repository *QuestionRepositoryMock) Find(ctx context.Context, DB *gorm.DB, param param.QuestionParam) ([]entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *QuestionRepositoryMock) FindOneBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *QuestionRepositoryMock) FindDetailed(ctx context.Context, DB *gorm.DB, param param.QuestionParam) ([]result.QuestionResult, error) {
	args := repository.Mock.Called(
		ctx, DB, param,
	)
	args0 := args.Get(0)
	args1 := args.Get(1)
	if args1 == nil {
		return args0.([]result.QuestionResult), nil
	}
	return args0.([]result.QuestionResult), args1.(error)
}

func (repository *QuestionRepositoryMock) FindOneDetailedBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (result.QuestionResult, error) {
	args := repository.Mock.Called(
		ctx, DB, param,
	)
	args0 := args.Get(0)
	args1 := args.Get(1)
	if args1 == nil {
		return args0.(result.QuestionResult), nil
	}
	return args0.(result.QuestionResult), args1.(error)
}

func (repository *QuestionRepositoryMock) FindDetailedLikedByUser(ctx context.Context, DB *gorm.DB, param param.LikedQuestionParam) ([]result.QuestionResult, error) {
	return []result.QuestionResult{}, nil
}
