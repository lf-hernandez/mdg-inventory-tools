package main

import (
	"fmt"
	"log"
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
	if item.PartNumber == "" {
		return fmt.Errorf("part number is required")
	}

	if item.Price == nil {
		return fmt.Errorf("price is required")
	}

	if item.Price != nil && *item.Price < 0 {
		return fmt.Errorf("price must be non-negative")
	}

	if item.Quantity == nil {
		return fmt.Errorf("quantity is required")
	}

	return nil
}

func logError(err error) {
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func logInfo(message string) {
	log.Printf("[INFO] %s", message)
}
