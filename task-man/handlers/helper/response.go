package helper

import (
	"encoding/json"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers/error"
	"net/http"
)

const (
	JsonType       = "application/json"
	DefaultCharset = "utf-8"
)

func SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		//TODO log
		statusCode = http.StatusInternalServerError
	}
	contentType := fmt.Sprintf("%s; carset=%s", JsonType, DefaultCharset)
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	if err != nil {
		//TODO log
	}
}

func SendErrorResponse(w http.ResponseWriter, err error.HttpError) {
	SendResponse(w, err.GetStatusCode(), err)
}
