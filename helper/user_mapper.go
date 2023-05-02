package helper

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/response"
)

func UserEntityToResponse(
	entity entity.User,
) response.UserResponse {
	return response.UserResponse{
		ID:       entity.ID,
		Username: entity.Username,
	}
}
