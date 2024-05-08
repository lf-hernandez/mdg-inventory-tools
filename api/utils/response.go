package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data any, headers http.Header) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	w.Header().Set("Content-Security-Policy", "frame-ancestors 'none';default-src 'self';")
	w.WriteHeader(status)
	w.Write(b)

	return nil
}
