package category

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
