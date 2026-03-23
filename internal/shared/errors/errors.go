package errors

import (
	"net/http"
)

// AppError represents a structured API error.
type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Predefined error constants
var (
	ErrNotFound     = &AppError{Status: http.StatusNotFound, Code: "NOT_FOUND", Message: "resource not found"}
	ErrUnauthorized = &AppError{Status: http.StatusUnauthorized, Code: "UNAUTHORIZED", Message: "authentication required"}
	ErrBadRequest   = &AppError{Status: http.StatusBadRequest, Code: "BAD_REQUEST", Message: "invalid request"}
	ErrForbidden    = &AppError{Status: http.StatusForbidden, Code: "FORBIDDEN", Message: "access denied"}
	ErrInternal     = &AppError{Status: http.StatusInternalServerError, Code: "INTERNAL", Message: "an unexpected error occurred"}
)

// NewAppError allows creating a custom AppError on the fly
func NewAppError(status int, code string, message string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
