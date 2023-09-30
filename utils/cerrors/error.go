package cerrors

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorType int

const (
	Unknown ErrorType = iota
	BadRequest
	NotFound
	InternalServerError
)

type CError struct {
	errorType     ErrorType
	originalError error
}

func (et ErrorType) New(message string) error {
	return &CError{
		errorType:     et,
		originalError: errors.New(message),
	}
}

func (et ErrorType) Wrap(err error, message string) error {
	return &CError{
		errorType:     et,
		originalError: fmt.Errorf("%s: %w", message, err),
	}
}

var _ error = (*CError)(nil)

func (e *CError) Error() string {
	return e.originalError.Error()
}

func StatusCode(err error) int {
	switch getType(err) {
	case BadRequest:
		return http.StatusBadRequest // 400
	case NotFound:
		return http.StatusNotFound // 404
	default:
		return http.StatusInternalServerError // 500
	}
}

func getType(err error) ErrorType {
	if ce, ok := err.(*CError); ok {
		return ce.errorType
	}
	return Unknown
}
