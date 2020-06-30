package helper

import (
	"encoding/json"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers/error"
	"go.uber.org/zap"
	"net/http"
)

const (
	JsonType       = "application/json"
	DefaultCharset = "utf-8"
)

type ResponseSender struct {
	logger *zap.Logger
}

func NewResponseSender(logger *zap.Logger) ResponseSender {
	return ResponseSender{logger: logger}
}

func (rs ResponseSender) SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	defer rs.logger.Sync()
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		rs.logger.Error(err.Error())
		statusCode = http.StatusInternalServerError
	}
	contentType := fmt.Sprintf("%s; carset=%s", JsonType, DefaultCharset)
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	if err != nil {
		rs.logger.Error(err.Error())
	}
}

func (rs ResponseSender) SendErrorResponse(w http.ResponseWriter, err error.HttpError) {
	rs.SendResponse(w, err.GetStatusCode(), err)
}
