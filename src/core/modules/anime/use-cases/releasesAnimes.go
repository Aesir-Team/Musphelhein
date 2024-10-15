package anime

import (
	"api/src/common/config"
	"api/src/core/modules/anime/types"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

func HomeAnime() string {
	cfg := config.GetConfig().Otakus
	return cfg
}

func scrapeAnimeReleases(url string) ([]types.Anime, error) {
	crawler := colly.NewCollector()
	var lancamentos []types.Anime 

	crawler.OnHTML("div.ultEpsContainer a", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.DOM.Find("img").AttrOr("alt", ""))
		link := e.Attr("href")

		imageLink := e.ChildAttr("img", "data-lazy-src")
		if imageLink == "" {
			imageLink = e.ChildAttr("noscript img", "src")
		}

		if title == "" {
			title = strings.TrimSpace(e.DOM.Text())
		}

		title = strings.ReplaceAll(title, "\n", "")
		title = strings.ReplaceAll(title, "\t", "")
		title = strings.TrimSpace(title)

		if title != "" && link != "" {
			anime := types.Anime{
				Title:     title,
				Link:      link,
				ImageLink: imageLink,
			}

			if strings.HasPrefix(imageLink, "https://otakuanimesscc.com/animes/capas/") {
				lancamentos = append(lancamentos, anime)
			}
		}
	})

	if err := crawler.Visit(url); err != nil {
		return nil, err
	}
	return lancamentos, nil
}

func ReleasesAnimeController(c *fiber.Ctx) error {
	url := ReleaseAnimes()
	lancamentos, err := scrapeAnimeReleases(url)
	if err != nil {
		log.Println("Error visiting URL:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to visit URL: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"lancamentos": lancamentos,
	})
}
