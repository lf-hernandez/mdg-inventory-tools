package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	ID          int     `json:"id"`
	ExternalID  string  `json:"external_id"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Quantity    float64 `json:"quantity"`
}

func main() {
	fmt.Printf("Starting server on port 8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "status: ok")
	})

	http.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleGet(w, r)
		case "POST":
		case "PUT":
		case "DELETE":
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Item{ID: 1, ExternalID: "001", Description: "Sample Data", Price: 10, Quantity: 2})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
