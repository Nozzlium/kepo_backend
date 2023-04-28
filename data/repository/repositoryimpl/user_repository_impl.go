package repositoryimpl

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/helper"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, user entity.User) entity.User {
	result := DB.WithContext(ctx).Create(&user)
	helper.PanicIfError(result.Error)
	return user
}

func (repository *UserRepositoryImpl) FindOneBy(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	result := entity.User{}
	find := DB.WithContext(ctx).Where(&user).First(&result)
	return result, find.Error
}
