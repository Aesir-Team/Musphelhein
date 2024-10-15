package anime

import (
	anime "api/src/core/modules/anime/use-cases"

	"github.com/gofiber/fiber/v2"
)

func AnimeController(app *fiber.App) {
	app.Get("/releases", anime.ReleasesAnimeController)
	app.Get("/latest-episodes", anime.LatestEpisodesController)
}
