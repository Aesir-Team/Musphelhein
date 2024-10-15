package anime

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

type AnimeReleases struct {
	Title         string `json:"title"`
	Link          string `json:"link"`
	ImageLink     string `json:"image_link"`
	EpisodeTitle  string `json:"episode_title"`
	EpisodeNumber string `json:"episode_number"`
	Quality       string `json:"quality"`
}

func ReleaseAnimes() string {
	return "https://otakuanimesscc.com/"
}

func scrapeAnimeLatests(url string) ([]AnimeReleases, error) {
	crawler := colly.NewCollector()

	crawler.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	var lancamentos []AnimeReleases

	crawler.OnHTML("div.ultEps div.ultEpsContainer div.ultEpsContainerItem a", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.ChildAttr("img", "alt"))
		episodetitle := strings.TrimSpace(e.DOM.Find("div.epInfos div.epNome").Text())
		epnumber := strings.TrimSpace(e.DOM.Find("div.epInfos div.epNum").Text())
		quality := strings.TrimSpace(e.DOM.Find("div.button-hd").Text())
		link := e.Attr("href")
		imageLink := e.ChildAttr("img", "data-lazy-src")

		if imageLink == "" {
			imageLink = e.ChildAttr("noscript img", "src")
		}

		title = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(title, "\n", ""), "\t", ""))

		if title != "" && link != "" {
			anime := AnimeReleases{
				Title:         title,
				Link:          link,
				ImageLink:     imageLink,
				EpisodeTitle:  episodetitle,
				EpisodeNumber: epnumber,
				Quality:       quality,
			}

			if strings.HasPrefix(imageLink, "https://otakuanimesscc.com/animes/images/") {
				lancamentos = append(lancamentos, anime)
			}
		}
	})

	if err := crawler.Visit(url); err != nil {
		return nil, err
	}

	return lancamentos, nil
}

func LatestEpisodesController(c *fiber.Ctx) error {
	url := ReleaseAnimes()
	lancamentos, err := scrapeAnimeLatests(url)
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
