package utils

import (
	"fmt"
	"strings"
)

func ExtractPathParam(path string, routePrefix string) (string, error) {
	param := strings.TrimPrefix(path, routePrefix)
	if param == "" || param == "/" {
		return "", fmt.Errorf("parameter is required")
	}

	return param, nil
}
