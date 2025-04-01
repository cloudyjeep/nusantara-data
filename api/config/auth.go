package config

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Auth struct{}

// validate token
func (a *Auth) Validate(token string) bool {
	// dummy token
	dataToken := []string{
		"abcd",
		"defg",
		"hijk",
		"lmno",
	}

	return slices.Contains(dataToken, token)
}

// midleware for enpoint protected
func (a *Auth) On(handler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// if token empty
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return LoadRequest(c).ReturnError(fiber.StatusUnauthorized, "Unauthorized", nil)
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// if token is valid -> next to handler
		if a.Validate(token) {
			return handler(c)
		}

		// if token invalid
		return LoadRequest(c).ReturnError(fiber.StatusUnauthorized, "Invalid token", nil)
	}
}
