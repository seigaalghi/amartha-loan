package errs

import (
	"fmt"
	"net/http"
)

type ErrorCode string

type CustomError struct {
	Code    int
	Message string
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%v: %s", e.Code, e.Message)
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func NewCustomErrorWithCause(code int, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func NewGeneralError(err error) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Err:     err,
	}
}
