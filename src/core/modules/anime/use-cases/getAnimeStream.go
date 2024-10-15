package anime

import (
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

func GetAnimeStream() string {
	return "https://otakuanimesscc.com/45279-9"
}

type Episode struct {
	Number string `json:"number"`
	Link   string `json:"link"`
}

type EpisodeDetails struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AudioType   string `json:"audioType"`
	ReleaseDate string `json:"releaseDate"`
	StreamURL   string `json:"streamUrl"`
}

func GetAnimeStreamController(c *fiber.Ctx) error {
	animeURL := c.Query("url")
	if animeURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL parameter is required.",
		})
	}

	crawler := colly.NewCollector()

	var playerURL string
	var episodes []Episode 
	var episodeDetails []EpisodeDetails 

	crawler.OnHTML("#player_1 iframe", func(e *colly.HTMLElement) {
		playerURL = e.Attr("src")
	})

	crawler.OnHTML(".EpsList a", func(e *colly.HTMLElement) {
		episodeNumber := e.Text 
		episodeLink := e.Attr("href")

		episode := Episode{
			Number: episodeNumber,
			Link:   episodeLink,
		}
		episodes = append(episodes, episode) 
	})

	crawler.OnHTML(".informacoes_ep_container .info", func(e *colly.HTMLElement) {
		infoText := strings.TrimSpace(e.Text)
		detail := EpisodeDetails{}
	
		switch {
		case e.Index == 0:
			if len(episodeDetails) < len(episodes) {
				detail.Title = infoText
				episodeDetails = append(episodeDetails, detail)
			}
		case e.Index == 1:
			if len(episodeDetails) > 0 {
				episodeDetails[len(episodeDetails)-1].Title += " - " + infoText 
			}
		}
	})

	crawler.OnError(func(r *colly.Response, err error) {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch the page.",
		})
	})

	err := crawler.Visit(animeURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to visit the URL.",
		})
	}

	if playerURL == "" {
		return c.JSON(fiber.Map{
			"error": "Player URL not found",
		})
	}

	response := fiber.Map{
		"playerURL":     playerURL,
		"episodes":      episodes,        
		"episodeDetails": episodeDetails,
	}

	return c.JSON(response)
}
