package user

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type UserService interface {
	FindOneBy(ctx context.Context, param param.UserParam) (response.UserResponse, error)
}
