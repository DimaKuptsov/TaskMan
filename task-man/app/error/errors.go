package error

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (err ValidationError) GetField() string {
	return err.Field
}

func (err ValidationError) GetMessage() string {
	return err.Message
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("\"%s\" field validation error occurred: %s", err.Field, err.Message)
}
