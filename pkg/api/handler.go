package api

import (
	"cipher-osint/pkg/ai"
	"cipher-osint/pkg/scraper"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func HandleIntelligence(c *fiber.Ctx) error {
	target := c.FormValue("target")
	model := c.FormValue("model")
	isGhost := c.FormValue("isGhost") == "true"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var rawIntel string
	var err error

	if isGhost {
		onionURL := fmt.Sprintf("http://jntlesuoc7kyjpsf.onion/index.php?q=%s", target)
		rawIntel, err = scraper.GhostScrape(onionURL)
		if err != nil {
			rawIntel = "Tor node link failed: falling back to surface recon."
		}
	} else {
		rawIntel = scraper.RunScraper(target)
	}

	nexus := ai.Nexus(model)
	analysis, _ := nexus.Analyze(ctx, rawIntel)

	return c.JSON(fiber.Map{
		"intel":  analysis,
		"status": "ZENITH_SUCCESS",
	})
}