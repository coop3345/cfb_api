package util

import (
	"encoding/json"
	"regexp"
	"strings"
)

func MarshalToJSONString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func Trim_endpoint(endpoint string) string {
	if !CONFIG.RUN_PARAMS.GET_WEEKLY && CONFIG.RUN_PARAMS.GET_FULL_SEASON {
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

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
