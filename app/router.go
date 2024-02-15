package app

import (
	"net/http"
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

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	router.GET("/details", userController.GetDetails)
	router.GET("/user/:id/details", userController.GetById)

	router.POST("/question", questionController.Create)
	router.GET("/question", questionController.Get)
	router.GET("/question/:id", questionController.GetById)
	router.GET("/user/:id/question", questionController.GetByUser)
	router.GET("/user/:id/question/like", questionController.GetLikedByUser)
	router.DELETE("/question/:id", questionController.Delete)
	router.PUT("/question/:id", questionController.Update)

	router.POST("/answer", answerController.Create)
	router.GET("/answer", answerController.Find)
	router.GET("/answer/:id", answerController.FindById)
	router.GET("/user/:id/answer", answerController.FindByUser)
	router.GET("/question/:id/answer", answerController.FindByQuestion)
	router.DELETE("/answer/:id", answerController.Delete)
	router.PUT("/answer/:id", answerController.Update)

	router.GET("/notification", notificationController.Find)
	router.PUT("/notification/:id/read", notificationController.Read)
	router.GET("/notification/unread", notificationController.GetUnreadCount)

	router.POST("/answer/like", answerLikeController.Like)
	router.POST("/question/like", questionLikeController.Like)

	router.GET("/category", categoryController.Get)

	return router

}
