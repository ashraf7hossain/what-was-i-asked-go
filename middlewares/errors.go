package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware for centralized error handling
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		// Retrieve any errors from the context
		errs := c.Errors
		if len(errs) > 0 {
			// Log the first error
			err := errs[0]
			log.Printf("Error occurred: %v", err.Err)

			// Determine status code
			statusCode := http.StatusInternalServerError
			if err.Meta != nil {
				if code, ok := err.Meta.(int); ok {
					statusCode = code
				}
			}

			// Default messages based on status code
			defaultMessages := map[int]string{
				http.StatusInternalServerError: "An internal server error occurred",
				http.StatusBadRequest:          "Invalid request",
				http.StatusUnauthorized:        "Unauthorized access",
				http.StatusForbidden:           "Access denied",
				http.StatusNotFound:            "Resource not found",
			}

			// Error response
			message := err.Error()
			if message == "" {
				message = defaultMessages[statusCode]
			}

			// Respond with error
			c.JSON(statusCode, gin.H{
				"message": defaultMessages[statusCode],
				"error":   message,
				"details": err.Error(),
			})
			return
		}
	}
}
