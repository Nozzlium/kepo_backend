package repositoryimpl

import (
	"context"
	"nozzlium/kepo_backend/data/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	result := DB.WithContext(ctx).Create(&user)
	return user, result.Error
}

func (repository *UserRepositoryImpl) FindOneBy(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	result := entity.User{}
	find := DB.WithContext(ctx).Where(&user).First(&result)
	return result, find.Error
}

func (repository *UserRepositoryImpl) FindOneBasedOnIdentity(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {
	result := entity.User{}
	find := DB.WithContext(ctx).Debug().Where(
		"username = ?", user.Username,
	).Or(
		"email = ?", user.Email,
	).First(&result)
	return result, find.Error
}
