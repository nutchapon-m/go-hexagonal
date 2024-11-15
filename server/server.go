package server

import (
	"fmt"
	v1 "go-hexagonal/server/routes/v1"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type server struct {
	app *fiber.App
}

func NewServ(app *fiber.App) server {
	app.Use(
		newLoggerMiddleWare(),
		newCorsMiddleWare(),
	)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	apiRoute := app.Group("/api/v1")
	v1.AuthRouter(apiRoute)

	return server{app: app}
}

func (s server) ListAndServ() {
	// Gracefully shutting down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		serv := <-c
		if serv.String() == "interrupt" {
			fmt.Println("Gracefully shutting down...")
			s.app.Shutdown()
		}
	}()

	if viper.GetString("server.mode") == "debug" {
		s.app.Listen("localhost:8000")
	} else {
		s.app.Listen(":" + viper.GetString("server.port"))
	}
}
