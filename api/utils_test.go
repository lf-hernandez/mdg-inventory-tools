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
