package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware for centralized error handling
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		// Retrieve all errors from the context
		errs := c.Errors
		if len(errs) > 0 {
			// Log all errors
			for _, err := range errs {
				log.Printf("Error occurred: %v", err.Err)
			}

			// Collect error details
			var errorDetails []gin.H
			for _, err := range errs {
				statusCode := http.StatusInternalServerError
				if err.Meta != nil {
					if code, ok := err.Meta.(int); ok {
						statusCode = code
					}
				}

				defaultMessages := map[int]string{
					http.StatusInternalServerError: "An internal server error occurred",
					http.StatusBadRequest:          "Invalid request",
					http.StatusUnauthorized:        "Unauthorized access",
					http.StatusForbidden:           "Access denied",
					http.StatusNotFound:            "Resource not found",
				}

				message := err.Error()
				if message == "" {
					message = defaultMessages[statusCode]
				}

				errorDetails = append(errorDetails, gin.H{
					"status_code": statusCode,
					"message":     message,
				})
			}

			// Determine the highest severity error code for the response
			finalStatusCode := errorDetails[0]["status_code"].(int)

			// Respond with all errors
			c.JSON(finalStatusCode, gin.H{
				"errors": errorDetails,
			})
			return
		}
	}
}
