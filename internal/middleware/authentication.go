package middleware

import (
	"context"
	"net/http"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhamadnirwansyah/authentication-service/domain"
	"github.com/muhamadnirwansyah/authentication-service/dto"
)

func Authenticate(authService domain.AuthenticationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.Split(c.Get("Authorization"), " ")
		if len(token) < 2 {
			return c.Status(http.StatusUnauthorized).JSON(dto.NewResponseMessage("Sorry, Token is invalid !"))
		}
		account, err := authService.Validate(context.Background(), token[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(dto.NewResponseMessage("Sorry, token is invalid !"))
		}
		c.Locals("x-account", account)
		return c.Next()
	}
}
