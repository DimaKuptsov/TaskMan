package logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	logger := GetLogger()
	if logger != nil {
		t.Errorf("Init: expected nil logger before init. Got %v", logger)
	}
	err := Init()
	if err != nil {
		t.Errorf("Init: expected initialization without error. Got %s", err.Error())
	}
	logger = GetLogger()
	if logger == nil {
		t.Errorf("Init: expected logger struct after init. Got nil")
	}
}
