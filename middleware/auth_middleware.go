package middleware

import (
	"context"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/exception"
	"nozzlium/kepo_backend/helper"
	"nozzlium/kepo_backend/tools"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(
	handle httprouter.Handle,
) AuthMiddleware {
	return AuthMiddleware{Handler: handle}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	auth := r.Header.Get("Authorization")
	token := strings.Split(auth, " ")
	if token[0] != "Bearer" {
		helper.PanicIfError(exception.BadRequestError{})
	}
	claims, err := tools.ParseJWTToken(token[1])
	if err != nil {
		panic(exception.UnauthorizedError{})
	}
	ctx := context.WithValue(r.Context(), constants.USER_ID_CLAIMS, claims)
	middleware.Handler(w, r.WithContext(ctx), params)
}
