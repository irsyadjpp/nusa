package error

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Internal   error  `json:"-"`
	StatusCode int    `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Internal)
	}
	return e.Message
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewWithStatus(code int, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewInternal(err error, message string) *AppError {
	return &AppError{
		Code:       500,
		Message:    message,
		Internal:   err,
		StatusCode: http.StatusInternalServerError,
	}
}

func BadRequest(message string) *AppError {
	return &AppError{
		Code:       400,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func Unauthorized(message string) *AppError {
	return &AppError{
		Code:       401,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func Forbidden(message string) *AppError {
	return &AppError{
		Code:       403,
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

func NotFound(message string) *AppError {
	return &AppError{
		Code:       404,
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func Conflict(message string) *AppError {
	return &AppError{
		Code:       409,
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

func ValidationError(message string) *AppError {
	return &AppError{
		Code:       422,
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

const (
	ErrCodeValidation    = "VALIDATION_ERROR"
	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeUnauthorized  = "UNAUTHORIZED"
	ErrCodeForbidden     = "FORBIDDEN"
	ErrCodeConflict      = "CONFLICT"
	ErrCodeInternal      = "INTERNAL_ERROR"
	ErrCodeDatabase      = "DATABASE_ERROR"
	ErrCodeExternalAPI   = "EXTERNAL_API_ERROR"
)
