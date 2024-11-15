package main

import (
	"go-hexagonal/internal/configs"
	"go-hexagonal/server"

	"github.com/gofiber/fiber/v2"
)

func init() {
	configs.Init()
	// db.InitPOSTGRESDB()
}

func main() {
	serv := server.NewServ(fiber.New())
	serv.ListAndServ()
}
