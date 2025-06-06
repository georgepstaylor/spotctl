package errors

import (
	"fmt"
	"net/http"
)

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeAPI represents API-related errors
	ErrorTypeAPI ErrorType = "API"
	// ErrorTypeConfig represents configuration-related errors
	ErrorTypeConfig ErrorType = "Config"
	// ErrorTypeValidation represents validation-related errors
	ErrorTypeValidation ErrorType = "Validation"
	// ErrorTypeInternal represents internal errors
	ErrorTypeInternal ErrorType = "Internal"
)

// Error represents a standardized error in the application
type Error struct {
	Type    ErrorType
	Code    int
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s error: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s error: %s", e.Type, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new API error
func NewAPIError(code int, message string, err error) *Error {
	return &Error{
		Type:    ErrorTypeAPI,
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NewConfigError creates a new configuration error
func NewConfigError(message string, err error) *Error {
	return &Error{
		Type:    ErrorTypeConfig,
		Code:    http.StatusInternalServerError,
		Message: message,
		Err:     err,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string, err error) *Error {
	return &Error{
		Type:    ErrorTypeValidation,
		Code:    http.StatusBadRequest,
		Message: message,
		Err:     err,
	}
}

// NewInternalError creates a new internal error
func NewInternalError(message string, err error) *Error {
	return &Error{
		Type:    ErrorTypeInternal,
		Code:    http.StatusInternalServerError,
		Message: message,
		Err:     err,
	}
}
