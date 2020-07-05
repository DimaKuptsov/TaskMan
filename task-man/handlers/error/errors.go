package error

import (
	"errors"
	"net/http"
)

const (
	BadRequestMessage              = "Bad Request"
	UnprocessableEntityMessage     = "Unprocessable Entity"
	InternalServerErrorMessage     = "Internal Server Error"
	InternalServerErrorDescription = "oops something went wrong, we are already solving this problem"
)

func NewBadRequestError(err error) HttpError {
	return NewHttpError(http.StatusBadRequest, BadRequestMessage, err)
}

func NewUnprocessableEntityError(err error) HttpError {
	return NewHttpError(http.StatusUnprocessableEntity, UnprocessableEntityMessage, err)
}

func NewInternalServerError(err error) HttpError {
	err = errors.New(InternalServerErrorDescription)
	return NewHttpError(http.StatusInternalServerError, InternalServerErrorMessage, err)
}
