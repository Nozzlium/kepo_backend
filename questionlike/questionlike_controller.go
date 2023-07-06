package questionlike

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QuestionLikeController interface {
	Like(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
