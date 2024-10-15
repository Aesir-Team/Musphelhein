package anime

import (
	"api/src/core/modules/anime/types"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

func scrapeAnimeEpisode(url string) (types.AnimeInfo, error) {
	crawler := colly.NewCollector()
	var animeInfo types.AnimeInfo

	crawler.OnHTML(".pageAnimeSection", func(e *colly.HTMLElement) {
		animeInfo.Title = strings.TrimSpace(e.ChildText("h1"))

		animeInfo.ImageLink = e.ChildAttr("div.animeCapa img", "data-lazy-src")
		if animeInfo.ImageLink == "" {
			animeInfo.ImageLink = e.ChildAttr("div.animeCapa img", "src")
		}

		e.ForEach(".animeInfos .animeInfo", func(_ int, elem *colly.HTMLElement) {
			switch {
			case strings.Contains(elem.Text, "Ano:"):
				animeInfo.Year = strings.TrimSpace(strings.Split(elem.Text, ":")[1])
			case strings.Contains(elem.Text, "Episódios:"):
				animeInfo.Episodes = strings.TrimSpace(strings.Split(elem.Text, ":")[1])
			case strings.Contains(elem.Text, "Audio:"):
				animeInfo.Audio = strings.TrimSpace(strings.Split(elem.Text, ":")[1])
			}
		})

		var genres []string
		e.ForEach(".animeGen li", func(_ int, elem *colly.HTMLElement) {
			genre := strings.TrimSpace(elem.Text)
			genres = append(genres, genre)
		})
		animeInfo.Genres = genres

		animeInfo.Synopsis = strings.TrimSpace(e.ChildText(".animeSecondContainer p"))

		e.ForEach(".sectionEpiInAnime a.list-epi", func(_ int, elem *colly.HTMLElement) {
			episode := types.EpisodeInfo{
				Title: strings.TrimSpace(elem.Attr("title")),
				Link:  elem.Attr("href"),
			}
			animeInfo.EpisodeList = append(animeInfo.EpisodeList, episode)
		})
	})

	err := crawler.Visit(url)
	if err != nil {
		return animeInfo, err
	}

	return animeInfo, nil
}
func AnimeInfoController(c *fiber.Ctx) error {
	animeURL := c.Query("url")

	animeInfo, err := scrapeAnimeEpisode(animeURL)
	if err != nil {
		log.Println("Error visiting URL:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to visit URL: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"animeInfo": animeInfo,
	})
}
