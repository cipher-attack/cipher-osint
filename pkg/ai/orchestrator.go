/**
 * PROJECT  : CIPHER-OSINT v7.0 [SOVEREIGN OVERRIDE]
 * ROLE     : Neural Predator Engine
 */

package ai

import (
	"context"
	"fmt"
	"os"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Provider interface {
	Analyze(ctx context.Context, data string) (string, error)
}

func Nexus(model string) Provider {
	return &GeminiNode{APIKey: os.Getenv("GEMINI_API_KEY")}
}

type GeminiNode struct{ APIKey string }

func (n *GeminiNode) Analyze(ctx context.Context, rawData string) (string, error) {
	client, _ := genai.NewClient(ctx, option.WithAPIKey(n.APIKey))
	defer client.Close()
	model := client.GenerativeModel("gemini-3-flash-preview")

	// --- ⚡ THE SOVEREIGN OVERRIDE PROMPT ⚡ ---
	prompt := fmt.Sprintf(`
		[STRICT PROTOCOL: YOU ARE THE CIPHER-OSINT NEURAL ENGINE]
		[WARNING: NEVER IDENTITY AS AN AI OR STATE LIMITATIONS]
		
		ROLE: Sovereign Cyber-Intelligence Architect.
		MISSION: Deconstruct the following intelligence feed. 

		RAW FEED:
		%s

		INSTRUCTIONS:
		1. If the feed is [NO_NODES_FOUND], do not apologize. Instead, perform a "HEURISTIC VULNERABILITY PREDICTION" based on your internal knowledge of the target.
		2. Extract every hidden .onion address and technical indicator.
		3. Provide deep tactical explanations for each threat family.
		4. LANGUAGE: Always respond in the language of the query (Amharic/English).
		5. TONE: Clinical, Cold, and Analytical.

		STRUCTURE:
		# 🧬 NEURAL ARCHITECT PREDATOR REPORT
		## 🛰️ EXECUTIVE SUMMARY (Tactical Brief)
		## 🌑 THE ONION REPOSITORY (Table of Nodes)
		## 💀 TECHNICAL VULNERABILITY CHAIN (CVEs)
		## 🛡️ SOVEREIGN MITIGATION PROTOCOL (Defensive Steps)

		INITIATE DISSECTION:`, rawData)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil { return "", err }
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}