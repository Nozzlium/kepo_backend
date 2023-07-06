package app

import (
	"nozzlium/kepo_backend/answer"
	"nozzlium/kepo_backend/answerlike"
	"nozzlium/kepo_backend/auth"
	"nozzlium/kepo_backend/question"
	"nozzlium/kepo_backend/questionlike"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	authController auth.AuthController,
	questionController question.QuestionController,
	questionLikeController questionlike.QuestionLikeController,
	answerController answer.AnswerController,
	answerlikeController answerlike.AnswerLikeController,
) *httprouter.Router {

	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)

	router.POST("/api/question", questionController.Create)
	router.GET("/api/question", questionController.Get)
	router.GET("/api/question/:id", questionController.GetById)
	router.POST("/api/question/like", questionLikeController.Like)

	router.POST("/api/answer", answerController.Create)
	router.GET("/api/answer", answerController.Find)
	router.GET("/api/answer/:id", answerController.FindById)
	router.GET("api/answer/like", answerlikeController.Like)

	return router

}
