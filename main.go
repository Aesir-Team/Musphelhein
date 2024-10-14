package main

import (
	"api/src/common/config"
	"api/src/common/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main () {
	app := fiber.New()

	handlers.RouteHandler(app)

	cfg := config.GetConfig()
	err := app.Listen(cfg.Port)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
}