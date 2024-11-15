package handlers

import (
	"go-hexagonal/internal/core/domains"
	"go-hexagonal/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type authHandlers struct {
	authSrv ports.AuthService
}

func NewAuthHandlers(authSrv ports.AuthService) ports.AuthHandler {
	return authHandlers{authSrv: authSrv}
}

func (h authHandlers) CreateUser(c *fiber.Ctx) error {
	user := domains.UserRequest{}

	if err := c.BodyParser(&user); err != nil {
		return handleError(c, err)
	}

	return c.SendString("user is created.")
}
