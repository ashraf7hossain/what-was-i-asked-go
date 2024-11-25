package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AbortWithError simplifies error handling with status codes
func AbortWithError(c *gin.Context, statusCode int, message string) {
	if message == "" {
		message = http.StatusText(statusCode)
	}
	c.Error(errors.New(message)).SetMeta(statusCode)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}

func NewError(message string) error {
	return gin.Error{Err: errors.New(message)}
}
