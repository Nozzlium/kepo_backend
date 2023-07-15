package main

import (
	"nozzlium/kepo_backend/app"
	"nozzlium/kepo_backend/auth"
	"nozzlium/kepo_backend/data/repository/repositoryimpl"
)

func main() {

	db := app.NewTestDB()

	userRepository := repositoryimpl.NewUserRepository()
	categoryRepository := repositoryimpl.NewCategoryRepository()
	questionRepository := repositoryimpl.NewQuestionRepository()
	questionLikeRepository := repositoryimpl.NewQuestionLikeRepository()
	answerRepository := repositoryimpl.NewAnswerRepository()
	answerLikeRepository := repositoryimpl.NewAnswerLikeRepository()

	authService := auth.NewAuthService(userRepository, db)

}
