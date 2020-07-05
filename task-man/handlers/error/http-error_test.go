package error

import (
	"errors"
	"github.com/DimaKuptsov/task-man/helpers"
	"strconv"
	"strings"
	"testing"
)

func TestNewHttpError(t *testing.T) {
	var tests = []struct {
		statusCode int
		message    string
		err        error
	}{
		{1, "test", errors.New("rand error")},
		{404, "not found", errors.New("error")},
		{500, InternalServerErrorMessage, errors.New(InternalServerErrorDescription)},
	}
	for _, test := range tests {
		httpErr := NewHttpError(test.statusCode, test.message, test.err)
		if httpErr.Code != test.statusCode {
			t.Errorf("NewHttpError: expected error with code %v. Got with code %v", test.statusCode, httpErr.Code)
		}
		if httpErr.Message != test.message {
			t.Errorf("NewHttpError: expected error with message %s. Got with message %s", test.message, httpErr.Message)
		}
		if httpErr.Description != test.err.Error() {
			t.Errorf("NewHttpError: expected error with description %s. Got with description %s", test.err.Error(), httpErr.Description)
		}
	}
}

func TestError(t *testing.T) {
	testErrors := getTestError()
	for _, testErr := range testErrors {
		errorStr := testErr.Error()
		if !strings.Contains(errorStr, strconv.Itoa(testErr.Code)) {
			t.Errorf("HttpError.Error: expected error string contain code '%v'. Error message %s", testErr.Code, errorStr)
		}
		if !strings.Contains(errorStr, testErr.Message) {
			t.Errorf("HttpError.Error: expected error string contain message '%s'. Error message %s", testErr.Message, errorStr)
		}
		if !strings.Contains(errorStr, testErr.Description) {
			t.Errorf("HttpError.Error: expected error string contain description '%s'. Error message %s", testErr.Description, errorStr)
		}
	}
}

func TestGetStatusCode(t *testing.T) {
	testErrors := getTestError()
	for _, testErr := range testErrors {
		if testErr.Code != testErr.GetStatusCode() {
			t.Errorf("HttpError.GetStatusCode: not equal code. From method %v. Code field %v", testErr.GetStatusCode(), testErr.Code)
		}
	}
}

func TestGetMessage(t *testing.T) {
	testErrors := getTestError()
	for _, testErr := range testErrors {
		if testErr.Message != testErr.GetMessage() {
			t.Errorf("HttpError.GetMessage: not equal message. From method %s. Message field %s", testErr.GetMessage(), testErr.Message)
		}
	}
}

func TestGetDescription(t *testing.T) {
	testErrors := getTestError()
	for _, testErr := range testErrors {
		if testErr.Description != testErr.GetDescription() {
			t.Errorf("HttpError.GetDescription: not equal description. From method %s. Description field %s", testErr.GetDescription(), testErr.Description)
		}
	}
}

func getTestError() []HttpError {
	testsNumber := 5
	var tests []HttpError
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		column := HttpError{
			Code:        helpers.GenerateIntBetween(0, 1000),
			Message:     helpers.GenerateRandomString(100, "string"),
			Description: helpers.GenerateRandomString(200, "error"),
		}
		tests = append(tests, column)
	}
	return tests
}
