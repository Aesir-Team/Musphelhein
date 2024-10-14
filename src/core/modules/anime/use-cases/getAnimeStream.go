package anime

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func GetAnimeStream() string {
	return "https://otakuanimesscc.com/45279-9"
}

func GetAnimeStreamController(c *fiber.Ctx) error {
	url := GetAnimeStream()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request.",
		})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch the page.",
		})
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response body.",
		})
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse the page.",
		})
	}

	iframe := doc.Find("#player_1 iframe") 
	playerURL, exists := iframe.Attr("src")
	if !exists {
		return c.JSON(fiber.Map{
			"error": "Player URL not found",
		})
	}

	return c.JSON(fiber.Map{
		"playerURL": playerURL,
	})
}
