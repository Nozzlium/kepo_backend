package category

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository"

	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: &mock.Mock{}}
var categoryService = CategoryServiceImpl{
	CategoryRepository: &categoryRepository,
}
var categoryController = CategoryControllerImpl{
	CategoryService: &categoryService,
}

var expectedCategories = []entity.Category{
	{
		ID:   1,
		Name: "cat1",
	},
	{
		ID:   2,
		Name: "cat2",
	},
}

func mockCategories() *mock.Call {
	mockCall := categoryRepository.Mock.On(
		"FindAll",
		mock.Anything,
		mock.Anything,
	).Return(
		expectedCategories,
		nil,
	)
	return mockCall
}
