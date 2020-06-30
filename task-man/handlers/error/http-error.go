package error

import "fmt"

type HttpError struct {
	Code        int    `json:"-"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewHttpError(statusCode int, message string, err error) HttpError {
	return HttpError{
		Code:        statusCode,
		Message:     message,
		Description: err.Error(),
	}
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Code, e.Message, e.Description)
}

func (e HttpError) GetStatusCode() int {
	return e.Code
}

func (e HttpError) GetMessage() string {
	return e.Message
}

func (e HttpError) GetDescription() string {
	return e.Description
}
