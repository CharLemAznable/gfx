package gclientx

import "fmt"

type HttpError interface {
	Error() string
	StatusCode() int
	StatusText() string
}

func NewHttpError(statusCode int, statusText string) error {
	return &localHttpError{statusCode: statusCode, statusText: statusText}
}

type localHttpError struct {
	statusCode int
	statusText string
}

func (e *localHttpError) Error() string {
	return fmt.Sprintf("%d %s", e.statusCode, e.statusText)
}

func (e *localHttpError) StatusCode() int {
	return e.statusCode
}

func (e *localHttpError) StatusText() string {
	return e.statusText
}
