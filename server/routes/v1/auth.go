package v1

import (
	"go-hexagonal/external/db"
	"go-hexagonal/internal/core/services"
	"go-hexagonal/internal/handlers"
	"go-hexagonal/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router) {
	userRepositories := repositories.NewUserRepositories(db.POSTGRESDB)
	authServices := services.NewAuthServices(userRepositories)
	authHandlers := handlers.NewAuthHandlers(authServices)

	router.Post("/user", authHandlers.CreateUser)

}
