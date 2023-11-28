package main

import "testing"

func Test_extractPathParam(t *testing.T) {
	result, err := extractPathParam("/api/path/prefix/123", "/api/path/prefix/")
	if err != nil {
		t.Error("errr occured:", err)
	}
	if result != "123" {
		t.Error("incorrect result: expected 123, got", result)
	}
}

func Test_errorHandlingExtractPathParam(t *testing.T) {
	_, err := extractPathParam("/", "/api/path/prefix/")
	if err.Error() != "parameter is required" {
		t.Error("incorrect result: expected parameter is required, got,", err)
	}

	_, err = extractPathParam("", "/api/path/prefix/")
	if err.Error() != "parameter is required" {
		t.Error("incorrect result: expected parameter is required, got,", err)
	}
}

func TestValidateItem(t *testing.T) {
	var (
		positivePrice float64 = 100.0
		negativePrice float64 = -50.0
		quantity      int     = 10
	)

	testCases := []struct {
		name    string
		item    Item
		wantErr bool
	}{
		{"Valid Item", Item{ID: "1", PartNumber: "PN1", Price: &positivePrice, Quantity: &quantity}, false},
		// {"Missing ID", Item{PartNumber: "PN1", Price: &positivePrice, Quantity: &quantity}, true},
		{"Negative Price", Item{ID: "1", PartNumber: "PN1", Price: &negativePrice, Quantity: &quantity}, true},
		{"Missing Quantity", Item{ID: "1", PartNumber: "PN1", Price: &positivePrice}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateItem(&tc.item)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateItem() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
