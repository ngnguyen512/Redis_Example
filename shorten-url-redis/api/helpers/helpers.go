package helpers

import (
	"os"
	"strings"
)

// EnforceHTTP ensures that a URL has an "http://" or "https://" prefix.
// If the prefix is missing, "http://" is added.
func EnforceHTTP(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}

// RemoveDomainError checks if the domain in the URL is the same as the environment variable "DOMAIN".
// It returns false if they are the same, meaning there's a "domain error".
func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	// Strip protocol and www. from URL for domain comparison
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	// Extract the domain part before the first slash
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}
