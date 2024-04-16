package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ExtractPathParam(path string, routePrefix string) (string, error) {
	param := strings.TrimPrefix(path, routePrefix)
	if param == "" || param == "/" {
		return "", fmt.Errorf("parameter is required")
	}

	return param, nil
}

func ExtractResourceFromURL(path string) (string, error) {
	re := regexp.MustCompile(`^/api/([^/]+)`)
	match := re.FindStringSubmatch(path)
	if len(match) >= 2 {
		return match[1], nil
	}
	return "", fmt.Errorf("Resource not found.")
}
