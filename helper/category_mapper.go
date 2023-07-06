package helper

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/response"
)

func CategoryEntityToResponse(
	entity entity.Category,
) response.CategoryResponse {
	return response.CategoryResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func CategoryEntityStructToResponseStruct(
	entities []entity.Category,
) []response.CategoryResponse {
	categories := []response.CategoryResponse{}
	for _, entity := range entities {
		categories = append(categories, CategoryEntityToResponse(entity))
	}
	return categories
}
