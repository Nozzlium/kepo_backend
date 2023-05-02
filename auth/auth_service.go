package auth

import (
	"context"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

type AuthService interface {
	Register(ctx context.Context, param param.AuthParam) (response.UserResponse, error)
	Login(ctx context.Context, param param.AuthParam) (response.AuthResponse, error)
}
