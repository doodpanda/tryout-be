package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware checks the claims for a user ID
func AuthMiddleware(c *fiber.Ctx) error {
	// Extract the token from the Authorization header
	userID := ""
	c.Context().SetUserValue("userID", userID)

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		// Allow access without a token
		return c.Next()
	}

	tokenString := strings.Split(authHeader, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		// Return the secret key
		return []byte("efUOEFHO8EFNKNmhnvsfjkhbOOIVADNIUubda"), nil
	})

	if err != nil {
		// Allow access without a valid token
		return c.Next()
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["uuid"]
		if userID != nil {
			// Store the user ID in the context
			c.Context().SetUserValue("userID", userID.(string))
		}
	} else {
		userID := ""
		c.Context().SetUserValue("userID", userID)
	}

	return c.Next()
}
