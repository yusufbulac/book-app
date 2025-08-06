package errorhandler

import "net/http"

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code string, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func NotFound(code, message string) *AppError {
	return NewAppError(code, message, http.StatusNotFound)
}

func BadRequest(code, message string) *AppError {
	return NewAppError(code, message, http.StatusBadRequest)
}

func InternalError(message string) *AppError {
	return NewAppError("INTERNAL_ERROR", message, http.StatusInternalServerError)
}
