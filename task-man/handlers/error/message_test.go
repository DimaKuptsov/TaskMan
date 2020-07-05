package error

import (
	"strings"
	"testing"
)

func TestGetBadParameterErrorMessage(t *testing.T) {
	var tests = []struct {
		parameter string
	}{
		{""},
		{"param"},
		{"random"},
	}
	for _, test := range tests {
		message := GetBadParameterErrorMessage(test.parameter)
		if len(message) == 0 {
			t.Error("GetBadParameterErrorMessage: expected not empty message. Got empty.")
		}
		if !strings.Contains(message, test.parameter) {
			t.Errorf("GetBadParameterErrorMessage: message must contain parameter '%s'. Message: %s", test.parameter, message)
		}
	}
}

func TestGetMissingParameterErrorMessage(t *testing.T) {
	var tests = []struct {
		parameter string
	}{
		{""},
		{"some"},
		{"parameter"},
	}
	for _, test := range tests {
		message := GetMissingParameterErrorMessage(test.parameter)
		if len(message) == 0 {
			t.Error("GetMissingParameterErrorMessage: expected not empty message. Got empty.")
		}
		if !strings.Contains(message, test.parameter) {
			t.Errorf("GetMissingParameterErrorMessage: message must contain parameter '%s'. Message: %s", test.parameter, message)
		}
	}
}
