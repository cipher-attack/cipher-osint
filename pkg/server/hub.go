package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"cipher-osint/pkg/ai"
	"cipher-osint/pkg/scraper"
)

func StartWebServer() {
	engine := html.New("./web/views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Title": "Cipher OSINT Dashboard"})
	})

	// API Endpoint for ai question
	app.Post("/api/analyze", func(c *fiber.Ctx) error {
		target := c.FormValue("target")
		raw := scraper.RunScraper(target)
		analysis, _ := ai.ProcessData(raw)
		return c.JSON(fiber.Map{"results": analysis})
	})

	app.Listen(":8080")
}