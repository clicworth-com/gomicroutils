package errs

import (
	"net/http"
)

type RestError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message,omitempty"`
}

func (e RestError) AsMessage() *RestError {
	return &RestError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func UnexpectedError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewUnexpectedError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func AuthenticationError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewValidationError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewAuthenticationError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}
