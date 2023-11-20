package main

import (
	"fmt"
	"strings"
)

func extractPathParam(path string, routePrefix string) (string, error) {
	param := strings.TrimPrefix(path, routePrefix)
	if param == "" || param == "/" {
		return "", fmt.Errorf("parameter is required")
	}

	return param, nil
}

func validateItem(item *Item) error {
	if item.ID == "" {
		return fmt.Errorf("ID is required")
	}

	if item.ExternalID == "" {
		return fmt.Errorf("external ID is required")
	}

	if item.Price == nil {
		return fmt.Errorf("price is requried")
	}

	if item.Price != nil && *item.Price < 0 {
		return fmt.Errorf("price must be non-negative")
	}

	if item.Quantity == nil {
		return fmt.Errorf("quantity is required")
	}

	return nil
}
