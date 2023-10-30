package answer

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AnswerController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
