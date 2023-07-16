package middleware

import (
	"context"
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/tools"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(
	handle http.Handler,
) *AuthMiddleware {
	return &AuthMiddleware{Handler: handle}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	auth := request.Header.Get("Authorization")
	token := strings.Split(auth, " ")
	ctx := request.Context()
	if token[0] == "Bearer" {
		claims, err := tools.ParseJWTToken(token[1])
		if err == nil {
			ctx = context.WithValue(ctx, constants.USER_ID_CLAIMS, *claims)
		}
	}
	middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
}
