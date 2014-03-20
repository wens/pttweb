package main

import (
	"net/http"
	"strings"
)

func checkCrawlerUserAgent(r *http.Request) bool {
	ua := r.UserAgent()

	switch {
	// Whitelist bots here
	case strings.Contains(ua, "Googlebot"):
		return true
	case strings.Contains(ua, "bingbot"):
		return true
	case strings.Contains(ua, "MSNbot"):
		return true
	case strings.Contains(ua, "facebookexternalhit"):
		return true
	}

	return false
}
