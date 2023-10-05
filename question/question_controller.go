package question

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QuestionController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetLikedByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
