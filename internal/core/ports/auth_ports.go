package ports

import (
	"go-hexagonal/internal/core/domains"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	CreateUser(user domains.UserRequest) error
}

type AuthHandler interface {
	CreateUser(c *fiber.Ctx) error
}
