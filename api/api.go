package main

import (
	"net/http"
	"strings"

	"github.com/lf-hernandez/mdg-inventory-tools/api/handlers"
)

func routeItems(deps *handlers.HandlerDependencies, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		deps.HandleGetItems(w, r)
	case http.MethodPost:
		deps.HandleCreateItem(w, r)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func routeSpecificItem(deps *handlers.HandlerDependencies, w http.ResponseWriter, r *http.Request) {
	itemID := strings.TrimPrefix(r.URL.Path, "/api/items/")
	if itemID == "" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		deps.HandleGetItem(w, r)
	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodPut:
		deps.HandleUpdateItem(w, r)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
