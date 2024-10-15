package anime

import (
	"api/src/common/config"
	"api/src/core/modules/anime/types"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

func SearchAnime() string {
	cfg := config.GetConfig().Otakus
	return cfg
}

func scrapeSearchAnime(url string) ([]types.Anime, error) {
	crawler := colly.NewCollector()
	var lancamentos []types.Anime

	crawler.OnHTML("div.loopAnimes div.ultAnisContainerItem a", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.DOM.Find("img").AttrOr("alt", ""))
		link := e.Attr("href")
		imageLink := e.ChildAttr("img", "src")

		if title != "" && link != "" {
			anime := types.Anime{
				Title:     title,
				Link:      link,
				ImageLink: imageLink,
			}
			lancamentos = append(lancamentos, anime)
		}
	})

	err := crawler.Visit(url)
	if err != nil {
		return nil, err
	}

	return lancamentos, nil
}

func SearchAnimeController(c *fiber.Ctx) error {
	url := SearchAnime()
	query := c.Query("q")
	page := c.Query("p")

	var lancamentos []types.Anime
	var err error

	if page == "" {
		lancamentos, err = scrapeSearchAnime(url + "?s=" + query)
	} else {
		lancamentos, err = scrapeSearchAnime(url + "/page/" + page + "?s=" + query)
	}

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

