package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetUserID retrieves the validated userID from the context
func GetUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get("userIDUint")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}

	return userID.(uint), nil
}
