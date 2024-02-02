package utils

import (
	"testing"

	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
)

func TestValidateItem(t *testing.T) {
	var (
		positivePrice float64 = 100.0
		negativePrice float64 = -50.0
		quantity      int     = 10
	)

	testCases := []struct {
		name    string
		item    models.Item
		wantErr bool
	}{
		{"Valid models.Item", models.Item{ID: "1", PartNumber: "PN1", Price: &positivePrice, Quantity: &quantity}, false},
		// {"Missing ID", models.Item{PartNumber: "PN1", Price: &positivePrice, Quantity: &quantity}, true},
		{"Negative Price", models.Item{ID: "1", PartNumber: "PN1", Price: &negativePrice, Quantity: &quantity}, true},
		{"Missing Quantity", models.Item{ID: "1", PartNumber: "PN1", Price: &positivePrice}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateItem(&tc.item)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateItem() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
