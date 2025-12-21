/**
 * PROJECT  : CIPHER-OSINT v7.0 [ZENITH OVERLORD]
 * ROLE     : Central Intelligence Orchestrator & Neural Hub
 * ARCHITECT: Biruk Getachew (CIPHER)
 * MOTTO    : "Deciphering the future before it happens.."
 */

package main

import (
	"cipher-osint/pkg/ai"
	"cipher-osint/pkg/scraper"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	// 1. Load Environment Variables
	_ = godotenv.Load()

	// 2. Ensure Intelligence Vaults Exist
	vaultPaths := []string{"vault/reports", "vault/system_logs", "web/static"}
	for _, path := range vaultPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			_ = os.MkdirAll(path, 0755)
		}
	}
}

func main() {
	// Initialize Sovereign Dashboard Engine
	app := fiber.New(fiber.Config{
		AppName: "CIPHER-OSINT Neural Interface v7.0",
		DisableStartupMessage: false,
	})

	// Enable Elite CORS Protocol
	app.Use(cors.New())

	// Static Asset Provisioning (CSS, JS, Images)
	app.Static("/static", "./web/static")

	// --- CORE ROUTES ---

	// Dashboard Endpoint: Serves the Zenith Interface
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/views/index.html")
	})

	// Intelligence API Node
	app.Post("/api/analyze", func(c *fiber.Ctx) error {
		target := c.FormValue("target")
		modelName := c.FormValue("model")
		isGhost := c.FormValue("isGhost") == "true"

		if target == "" {
			return c.Status(400).JSON(fiber.Map{"error": "TARGET_REQUIRED"})
		}

		fmt.Printf("\n[%s] 🛰️  NEW_RECON_TASK: Target: %s | Node: %s | Ghost: %v\n", 
			time.Now().Format("15:04:05"), target, modelName, isGhost)

		// Create High-Priority Context
		ctx, cancel := context.WithTimeout(context.Background(), 70*time.Second)
		defer cancel()

		// Step 1: Execute Multi-Source Intelligence Scrape
		surfaceIntel := scraper.RunScraper(target)
		darkIntel := scraper.SearchDarkWeb(target)
		
		combinedRawData := fmt.Sprintf(
			"[SURFACE_WEB_RECON]\n%s\n\n[DARK_WEB_GHOST_NODES]\n%v", 
			surfaceIntel, darkIntel,
		)

		// Step 2: Engage Neural Nexus
		analyzer := ai.Nexus(modelName)
		synthesis, err := analyzer.Analyze(ctx, combinedRawData)
		if err != nil {
			fmt.Printf("[!] NEURAL_LINK_FAILURE: %v\n", err)
			return c.Status(500).JSON(fiber.Map{
				"error":  "Neural Synthesis Error: " + err.Error(),
				"status": "CRITICAL_FAILURE",
			})
		}

		// Step 3: Sovereign Vaulting (Auto-Save Report)
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("INTEL_%s_%d.md", target, timestamp)
		reportPath := filepath.Join("vault/reports", fileName)
		_ = os.WriteFile(reportPath, []byte(synthesis), 0644)

		fmt.Printf("[+] [VAULTED] Intelligence Secured: %s\n", reportPath)

		// Step 4: Dispatch Analysis to Dashboard
		return c.JSON(fiber.Map{
			"results": synthesis,
			"vault":   reportPath,
			"status":  "CLOSINT_SUCCESS",
			"sid":     fmt.Sprintf("OVL-%d", timestamp),
		})
	})

	// 3. Launch Sovereign Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	banner := `
       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
  ▄▓▓▓▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
 ▄▓▓▓▀  CIPHER-OSINT v7.0
 ▓▓▓▀   ZENITH OVERLORD ONLINE
 ▓▓▓▀   PORT: %s
 ▀▓▓▓▄  STATUS: SOVEREIGN
   ▀▓▓▓▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
     ▀▀▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▀
	`
	fmt.Printf(banner, port)
	fmt.Printf("\n[+] Access Terminal: http://localhost:%s\n", port)
	fmt.Printf("[+] Deciphering the future before it happens..\n\n")

	log.Fatal(app.Listen(":" + port))
}