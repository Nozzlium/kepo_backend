package app

import (
	"nozzlium/kepo_backend/answer"
	"nozzlium/kepo_backend/answerlike"
	"nozzlium/kepo_backend/auth"
	"nozzlium/kepo_backend/category"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/question"
	"nozzlium/kepo_backend/questionlike"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {

	db := NewTestDB()
	validator := validator.New()

	userRepository := repository.NewUserRepository()
	categoryRepository := repository.NewCategoryRepository()
	questionRepository := repository.NewQuestionRepository()
	questionLikeRepository := repository.NewQuestionLikeRepository()
	answerRepository := repository.NewAnswerRepository()
	answerLikeRepository := repository.NewAnswerLikeRepository()

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

	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)

	router.POST("/api/question", questionController.Create)
	router.GET("/api/question", questionController.Get)
	router.GET("/api/question/:id", questionController.GetById)
	router.GET("/api/user/:id/question", questionController.GetByUser)

	router.POST("/api/answer", answerController.Create)
	router.GET("/api/answer", answerController.Find)
	router.GET("/api/answer/:id", answerController.FindById)
	router.GET("/api/user/:id/answer", answerController.FindByUser)
	router.GET("/api/question/:id/answer", answerController.FindByQuestion)

	router.POST("/api/answer/like", answerLikeController.Like)
	router.POST("/api/question/like", questionLikeController.Like)

	router.GET("/api/category", categoryController.Get)

	return router

}
