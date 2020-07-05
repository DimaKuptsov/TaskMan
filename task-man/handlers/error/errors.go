package error

import (
	"errors"
	"net/http"
)

func NewBadRequestError(err error) HttpError {
	return NewHttpError(http.StatusBadRequest, "Bad Request", err)
}

func NewUnprocessableEntityError(err error) HttpError {
	return NewHttpError(http.StatusUnprocessableEntity, "Unprocessable Entity", err)
}

func NewInternalServerError(err error) HttpError {
	err = errors.New("oops something went wrong, we are already solving this problem")
	return NewHttpError(http.StatusInternalServerError, "Internal Server Error", err)
}
