package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func SendRestError(c *gin.Context, err *RestError) {
	c.JSON(err.Status, err)
}

// 400 Bad Request
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// 401 Unauthorized
func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

// 403 Forbidden
func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "forbidden",
	}
}

// 404 Not Found
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// 409 Conflict
func NewConflictError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusConflict,
		Error:   "not_found",
	}
}

// 500 Internal Server Error
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
