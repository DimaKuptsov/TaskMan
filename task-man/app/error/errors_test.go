package error

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"strings"
	"testing"
)

func TestValidationErrorGetField(t *testing.T) {
	tests := getTestValidateError()
	for _, testErr := range tests {
		if testErr.GetField() != testErr.Field {
			t.Errorf("ValidationError.GetField: not equal field. From method %s. Field field %s", testErr.GetField(), testErr.Field)
		}
	}
}

func TestValidationErrorGetMessage(t *testing.T) {
	tests := getTestValidateError()
	for _, testErr := range tests {
		if testErr.GetMessage() != testErr.Message {
			t.Errorf("ValidationError.GetMessage: not equal message. From method %s. Message field %s", testErr.GetMessage(), testErr.Message)
		}
	}
}

func TestValidationErrorError(t *testing.T) {
	tests := getTestValidateError()
	for _, testErr := range tests {
		errorText := testErr.Error()
		if !strings.Contains(errorText, testErr.GetField()) {
			t.Errorf("ValidationError.Error: error text doesn't contain field. Error text %s. Field %s", errorText, testErr.GetField())
		}
		if !strings.Contains(errorText, testErr.GetMessage()) {
			t.Errorf("ValidationError.Error: error text doesn't contain message. Error text %s. Message %s", errorText, testErr.GetMessage())
		}
	}
}

func getTestValidateError() []ValidationError {
	testsNumber := 5
	var tests []ValidationError
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		randomField := helpers.GenerateRandomString(10, "field")
		randomMessage := helpers.GenerateRandomString(100, "message")
		validationError := ValidationError{Field: randomField, Message: randomMessage}
		tests = append(tests, validationError)
	}
	return tests
}
