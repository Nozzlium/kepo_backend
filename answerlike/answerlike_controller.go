package answerlike

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AnswerLikeController interface {
	Like(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
