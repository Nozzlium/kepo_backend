package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	GetDetails(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
