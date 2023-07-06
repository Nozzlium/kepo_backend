package category

import (
	"context"
	"nozzlium/kepo_backend/data/response"
)

type CategoryService interface {
	Get(ctx context.Context) ([]response.CategoryResponse, error)
}
