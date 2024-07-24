package main

import (
	"fmt"
	"go-hexagonal/api/middlewares"
	"go-hexagonal/api/routes"
	"go-hexagonal/configs"
	"go-hexagonal/pkg/db"
	"os"
	"os/signal"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func init() {
	configs.Init()
	db.InitDB()
}

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 10 * time.Minute,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middlewares.NewCorsMiddleWare())
	app.Use(middlewares.NewLoggerMiddleWare())

	r := app.Group("/v1/api")
	routes.User(r)

	// Gracefully shutting down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		s := <-c
		if s.String() == "interrupt" {
			fmt.Println("Gracefully shutting down...")
			app.Shutdown()
		}
	}()

	app.Listen("localhost:" + configs.GetString("SERVER_PORT"))
}
