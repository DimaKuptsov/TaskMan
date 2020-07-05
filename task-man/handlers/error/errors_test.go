package error

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	var tests = []struct {
		err error
	}{
		{errors.New("some error")},
		{errors.New("other error")},
		{errors.New("new error")},
	}
	for _, test := range tests {
		badRequestErr := NewBadRequestError(test.err)
		if badRequestErr.Code != http.StatusBadRequest {
			t.Errorf("NewBadRequestError: expected error with code %v. Got with code %v", http.StatusBadRequest, badRequestErr.Code)
		}
		if badRequestErr.Message != BadRequestMessage {
			t.Errorf("NewBadRequestError: expected error with message %s. Got with message %s", BadRequestMessage, badRequestErr.Message)
		}
		if badRequestErr.Description != test.err.Error() {
			t.Errorf("NewBadRequestError: expected error with description %s. Got with description %s", test.err.Error(), badRequestErr.Description)
		}
	}
}

func TestNewUnprocessableEntityError(t *testing.T) {
	var tests = []struct {
		err error
	}{
		{errors.New("rand error")},
		{errors.New("error")},
		{errors.New("new error message")},
	}
	for _, test := range tests {
		unprocessableEntityErr := NewUnprocessableEntityError(test.err)
		if unprocessableEntityErr.Code != http.StatusUnprocessableEntity {
			t.Errorf("NewUnprocessableEntityError: expected error with code %v. Got with code %v", http.StatusBadRequest, unprocessableEntityErr.Code)
		}
		if unprocessableEntityErr.Message != UnprocessableEntityMessage {
			t.Errorf("NewUnprocessableEntityError: expected error with message %s. Got with message %s", BadRequestMessage, unprocessableEntityErr.Message)
		}
		if unprocessableEntityErr.Description != test.err.Error() {
			t.Errorf("NewUnprocessableEntityError: expected error with description %s. Got with description %s", test.err.Error(), unprocessableEntityErr.Description)
		}
	}
}

func TestNewInternalServerError(t *testing.T) {
	var tests = []struct {
		err error
	}{
		{errors.New("rand error")},
		{errors.New("error")},
		{errors.New("new error message")},
	}
	for _, test := range tests {
		internalServerErr := NewInternalServerError(test.err)
		if internalServerErr.Code != http.StatusInternalServerError {
			t.Errorf("NewInternalServerError: expected error with code %v. Got with code %v", http.StatusBadRequest, internalServerErr.Code)
		}
		if internalServerErr.Message != InternalServerErrorMessage {
			t.Errorf("NewInternalServerError: expected error with message %s. Got with message %s", BadRequestMessage, internalServerErr.Message)
		}
		if internalServerErr.Description != InternalServerErrorDescription {
			t.Errorf("NewUnprocessableEntityError: expected error with description %s. Got with description %s", InternalServerErrorDescription, internalServerErr.Description)
		}
	}
}
