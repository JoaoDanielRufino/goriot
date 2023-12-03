package request

import "fmt"

type RequestError struct {
	Message    string
	StatusCode uint
}

func (e RequestError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.StatusCode, e.Message)
}
