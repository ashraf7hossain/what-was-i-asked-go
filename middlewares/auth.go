package middlewares

import (
	"fmt"
	"net/http"
	"rest-in-go/initializers"
	"rest-in-go/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Split the token from the "Bearer " prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Parse the token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token method is HMAC and check for validity
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// Return the JWT signing key from your environment variables or constants
			return []byte(initializers.SecretKey), nil
		})

		// Handle token parsing errors
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract user ID from the token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["userID"].(float64)
			c.Set("userID", uint(userID)) // Store the user ID in the context
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Continue processing the request
		c.Next()
	}
}

// ExtractUserIDMiddleware adds userID to the request context after validation
func ExtractUserIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			utils.AbortWithError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		userIDUint, ok := userID.(uint)
		if !ok {
			utils.AbortWithError(c, http.StatusBadRequest, "Invalid user ID format")
			return
		}

		c.Set("userIDUint", userIDUint) // Add validated userID to context
		c.Next()
	}
}