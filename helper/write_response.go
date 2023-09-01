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
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
