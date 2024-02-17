package app

import (
	"nozzlium/kepo_backend/answer"
	"nozzlium/kepo_backend/answerlike"
	"nozzlium/kepo_backend/auth"
	"nozzlium/kepo_backend/category"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/notification"
	"nozzlium/kepo_backend/question"
	"nozzlium/kepo_backend/questionlike"
	"nozzlium/kepo_backend/user"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {

	db := NewDB()
	validator := validator.New()

	userRepository := repository.NewUserRepository()
	categoryRepository := repository.NewCategoryRepository()
	questionRepository := repository.NewQuestionRepository()
	questionLikeRepository := repository.NewQuestionLikeRepository()
	answerRepository := repository.NewAnswerRepository()
	answerLikeRepository := repository.NewAnswerLikeRepository()
	notificationRepository := repository.NewNotificationRepository()

	authService := auth.NewAuthService(userRepository, db)
	categoryService := category.NewCategoryService(db, categoryRepository)
	questionService := question.NewQuestionService(questionRepository, db)
	questionLikeService := questionlike.NewQuestionLikeService(
		questionLikeRepository,
		questionRepository,
		userRepository,
		notificationRepository,
		db,
	)
	answerService := answer.NewAnswerService(
		answerRepository,
		notificationRepository,
		db,
	)
	answerLikeService := answerlike.NewAnswerLikeService(
		answerLikeRepository,
		answerRepository,
		userRepository,
		notificationRepository,
		db,
	)
	userService := user.NewUserService(userRepository, db)
	userController := user.NewUserController(userService)

	notificationService := notification.NewNotificationService(
		notificationRepository,
		db,
	)
	notificationController := notification.NewNotificationController(
		notificationService,
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

	router.GET("/api/details", userController.GetDetails)
	router.GET("/api/user/:id/details", userController.GetById)

	router.POST("/api/question", questionController.Create)
	router.GET("/api/question", questionController.Get)
	router.GET("/api/question/:id", questionController.GetById)
	router.GET("/api/user/:id/question", questionController.GetByUser)
	router.GET("/api/user/:id/question/like", questionController.GetLikedByUser)
	router.DELETE("/api/question/:id", questionController.Delete)
	router.PUT("/api/question/:id", questionController.Update)

	router.POST("/api/answer", answerController.Create)
	router.GET("/api/answer", answerController.Find)
	router.GET("/api/answer/:id", answerController.FindById)
	router.GET("/api/user/:id/answer", answerController.FindByUser)
	router.GET("/api/question/:id/answer", answerController.FindByQuestion)
	router.DELETE("/api/answer/:id", answerController.Delete)
	router.PUT("/api/answer/:id", answerController.Update)

	router.GET("/api/notification", notificationController.Find)
	router.PUT("/api/notification/:id/read", notificationController.Read)
	router.GET("/api/notification/unread", notificationController.GetUnreadCount)

	router.POST("/api/answer/like", answerLikeController.Like)
	router.POST("/api/question/like", questionLikeController.Like)

	router.GET("/api/category", categoryController.Get)

	return router

}
