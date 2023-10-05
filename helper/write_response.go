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
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTION")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
