package scraper

import (
	"context"
	"fmt"
	"net" //  new
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
)

const TorProxy = "127.0.0.1:9050"

func GhostScrape(onionURL string) (string, error) {
	dialer, err := proxy.SOCKS5("tcp", TorProxy, nil, proxy.Direct)
	if err != nil {
		return "", fmt.Errorf("tor_gateway_offline: %w", err)
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	client := &http.Client{Transport: transport, Timeout: 45 * time.Second}

	resp, err := client.Get(onionURL)
	if err != nil {
		return "", fmt.Errorf("dark_node_unreachable: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("document_parse_failure: %w", err)
	}

	doc.Find("script, style, iframe, img, footer, nav").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	cleanText := strings.Join(strings.Fields(doc.Text()), " ")
	if len(cleanText) > 4000 {
		cleanText = cleanText[:4000] + "... [INTEL_TRUNCATED]"
	}

	return cleanText, nil
}