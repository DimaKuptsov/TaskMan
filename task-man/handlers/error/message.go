package error

import "fmt"

func GetMissingParameterErrorMessage(parameter string) string {
	return fmt.Sprintf("missing required parameter '%s'", parameter)
}

func GetBadParameterErrorMessage(parameter string) string {
	return fmt.Sprintf("invalid parameter '%s'", parameter)
}
