package routes

import (
	"api/src/core/modules/anime"

	"github.com/gofiber/fiber/v2"
)


func MainRoute(app *fiber.App) {
	anime.AnimeController(app)
}