package gclientx

import "fmt"

type statusError struct {
	status  int
	message string
}

func (e *statusError) Error() string {
	return fmt.Sprintf("%d: %s", e.status, e.message)
}

func NewStatusError(status int, message string) error {
	return &statusError{status: status, message: message}
}
