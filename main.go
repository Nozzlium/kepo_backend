package main

import (
	"net/http"
	"nozzlium/kepo_backend/app"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/middleware"
)

func main() {
	router := app.NewRouter()

	router.PanicHandler = exception.ErrorHandler

	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:2637",
		Handler: authMiddleware,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
