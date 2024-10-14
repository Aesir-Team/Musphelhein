package handlers

import (
	"api/src/core/routes"

	"github.com/gofiber/fiber/v2"
)

func RouteHandler(app *fiber.App) {
	routes.MainRoute(app)
}