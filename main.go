package main

import (
	"net/http"
	"nozzlium/kepo_backend/answer"
	"nozzlium/kepo_backend/answerlike"
	"nozzlium/kepo_backend/app"
	"nozzlium/kepo_backend/auth"
	"nozzlium/kepo_backend/category"
	"nozzlium/kepo_backend/data/repository/repositoryimpl"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/middleware"
	"nozzlium/kepo_backend/question"
	"nozzlium/kepo_backend/questionlike"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewTestDB()
	validator := validator.New()

	userRepository := repositoryimpl.NewUserRepository()
	categoryRepository := repositoryimpl.NewCategoryRepository()
	questionRepository := repositoryimpl.NewQuestionRepository()
	questionLikeRepository := repositoryimpl.NewQuestionLikeRepository()
	answerRepository := repositoryimpl.NewAnswerRepository()
	answerLikeRepository := repositoryimpl.NewAnswerLikeRepository()

	authService := auth.NewAuthService(userRepository, db)
	categoryService := category.NewCategoryService(db, categoryRepository)
	questionService := question.NewQuestionService(questionRepository, db)
	questionLikeService := questionlike.NewQuestionLikeService(
		questionLikeRepository, questionRepository, db,
	)
	answerService := answer.NewAnswerService(answerRepository, db)
	answerLikeService := answerlike.NewAnswerLikeService(
		answerLikeRepository, answerRepository, db,
	)

	authController := auth.NewAuthController(authService, validator)
	categoryController := category.NewCategoryController(categoryService)
	questionController := question.NewQuestionController(questionService, validator)
	questionLikeController := questionlike.NewQuestionLikeController(
		questionLikeService, validator,
	)
	answerController := answer.NewAnswerController(answerService, validator)
	answerLikeController := answerlike.NewAnswerLikeController(
		answerLikeService, validator,
	)

	router := app.NewRouter(
		authController,
		questionController,
		questionLikeController,
		answerController,
		answerLikeController,
		categoryController,
	)

	router.PanicHandler = exception.ErrorHandler

	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:2637",
		Handler: authMiddleware,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
