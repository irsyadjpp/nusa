package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, Response{
			Success: false,
			Error: &Error{
				Code:    getCodeFromAppError(appErr),
				Message: appErr.Message,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error: &Error{
			Code:    ErrCodeInternal,
			Message: "Internal server error",
		},
	})
}

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func CreatedResponse(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func MessageResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
	})
}

func getCodeFromAppError(err *AppError) string {
	switch err.StatusCode {
	case http.StatusBadRequest:
		return ErrCodeValidation
	case http.StatusUnauthorized:
		return ErrCodeUnauthorized
	case http.StatusForbidden:
		return ErrCodeForbidden
	case http.StatusNotFound:
		return ErrCodeNotFound
	case http.StatusConflict:
		return ErrCodeConflict
	default:
		return ErrCodeInternal
	}
}
