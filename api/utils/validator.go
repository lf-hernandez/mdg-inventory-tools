package utils

import (
	"fmt"

	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
)

func ValidateItem(item *models.Item) error {
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
