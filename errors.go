package errors

import (
	"errors"
	"fmt"
	"runtime/debug"
)

const ERROR_MISSING_REQUIRED string = "error missing required values"

type AppError struct {
	Inner      error
	Message    string
	StackTrace string
}

func NewAppError(errMsg string, msgArgs ...interface{}) AppError {
	errMsg = fmt.Sprintf(errMsg, msgArgs...)
	return AppError{
		Inner:      errors.New(errMsg),
		Message:    errMsg,
		StackTrace: string(debug.Stack()),
	}
}

func WrapError(err error, msgf string, msgArgs ...interface{}) AppError {
	errMsg := err.Error()
	if errMsg != "" {
		errMsg = fmt.Sprintf(msgf, msgArgs...)
	}
	return AppError{
		Inner:      err,
		Message:    errMsg,
		StackTrace: string(debug.Stack()),
	}
}

func (err AppError) Error() string {
	return err.Message
}
