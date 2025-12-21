package scraper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func SearchDarkWeb(query string) []string {
	// "ransomeware"logic
	fullQuery := query + " onion link leak"
	searchURL := fmt.Sprintf("https://ahmia.fi/search/?q=%s", url.QueryEscape(fullQuery))

	client := &http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("GET", searchURL, nil)

	// real browser theft
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://ahmia.fi/")

	resp, err := client.Do(req)
	if err != nil { return []string{"[!] NODE_UNREACHABLE"} }
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)

	// regex
	re := regexp.MustCompile(`[a-z2-7]{16,56}\.onion`)
	matches := re.FindAllString(bodyStr, -1)

	var targets []string
	unique := make(map[string]bool)

	for _, m := range matches {
		m = strings.ToLower(m)
		// search bars and remove junk files
		if !strings.Contains(m, "ahmia") && !strings.Contains(m, "juhanurmi") && !unique[m] {
			unique[m] = true
			targets = append(targets, "http://"+m)
		}
	}

	if len(targets) == 0 { return []string{"[NO_NODES_FOUND]"} }
	return targets
}