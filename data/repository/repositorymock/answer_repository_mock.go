package repositorymock

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/result"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type AnswerRepositoryMock struct {
	Mock mock.Mock
}

func (repository *AnswerRepositoryMock) Insert(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error) {
	args := repository.Mock.Called(ctx, DB, answer)
	args0 := args.Get(0)
	args1 := args.Get(1)
	if args1 == nil {
		return args0.(entity.Answer), nil
	}
	return args0.(entity.Answer), args1.(error)
}

func (repository *AnswerRepositoryMock) FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]entity.Answer, error) {
	return []entity.Answer{}, nil
}

func (repository *AnswerRepositoryMock) FindOneBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (entity.Answer, error) {
	return entity.Answer{}, nil
}

func (repository *AnswerRepositoryMock) FindDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]result.AnswerResult, error) {
	args := repository.Mock.Called(ctx, DB, param)
	args0 := args.Get(0)
	args1 := args.Get(1)
	if args1 == nil {
		return args0.([]result.AnswerResult), nil
	}
	return args0.([]result.AnswerResult), args1.(error)
}
