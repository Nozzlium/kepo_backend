package helper

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(
	writer http.ResponseWriter,
	response interface{},
) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
