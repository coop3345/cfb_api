package util

import (
	"regexp"
	"strings"
)

func Trim_endpoint(endpoint string) string {
	if !GET_WEEKLY && GET_FULL_SEASON {
		re := regexp.MustCompile(`([?&])week=\d+(&)?`)

		// Replace with either '?' or '&' depending on what's left
		cleaned := re.ReplaceAllStringFunc(endpoint, func(m string) string {
			// If match ends with '&', we keep the leading character (usually '?')
			if strings.HasSuffix(m, "&") {
				return string(m[0])
			}
			// If it's the last param, just remove it
			return ""
		})

		// Clean up any trailing `&` or `?`
		cleaned = strings.TrimSuffix(cleaned, "&")
		cleaned = strings.TrimSuffix(cleaned, "?")

		return cleaned
	}

	return endpoint
}
