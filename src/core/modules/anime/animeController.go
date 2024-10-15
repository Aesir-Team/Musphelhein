package anime

import (
	anime "api/src/core/modules/anime/use-cases"

	"github.com/gofiber/fiber/v2"
)

func AnimeController(app *fiber.App) {
	app.Get("/releases", anime.ReleasesAnimeController)
	app.Get("/latest-episodes", anime.LatestEpisodesController)
	app.Get("/search-anime", anime.SearchAnimeController)
	app.Get("/anime-info", anime.AnimeInfoController)
	app.Get("/stream-anime", anime.GetAnimeStreamController)
}
