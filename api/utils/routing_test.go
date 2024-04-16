package utils

import "testing"

func Test_extractPathParam(t *testing.T) {
	result, err := ExtractPathParam("/api/path/prefix/123", "/api/path/prefix/")
	if err != nil {
		t.Error("errr occured:", err)
	}
	if result != "123" {
		t.Error("incorrect result: expected 123, got", result)
	}
}

func Test_errorHandlingExtractPathParam(t *testing.T) {
	_, err := ExtractPathParam("/", "/api/path/prefix/")
	if err.Error() != "parameter is required" {
		t.Error("incorrect result: expected parameter is required, got,", err)
	}

	_, err = ExtractPathParam("", "/api/path/prefix/")
	if err.Error() != "parameter is required" {
		t.Error("incorrect result: expected parameter is required, got,", err)
	}
}

func TestExtractResourceFromURL(t *testing.T) {
	tests := []struct {
		path           string
		expectedResult string
		expectedError  bool
	}{
		{"/api/items/123", "items", false},
		{"/api/users", "users", false},
		{"/api", "", true},
		{"/", "", true},
	}

	for _, test := range tests {
		result, err := ExtractResourceFromURL(test.path)
		if test.expectedError && err == nil {
			t.Errorf("Expected an error but got none for path: %s", test.path)
		} else if !test.expectedError && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != test.expectedResult {
			t.Errorf("Expected result %s but got %s for path: %s", test.expectedResult, result, test.path)
		}
	}
}
