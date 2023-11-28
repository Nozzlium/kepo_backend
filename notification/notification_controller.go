package notification

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NotificationController interface {
	Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Read(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
