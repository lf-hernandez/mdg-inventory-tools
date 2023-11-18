package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type Item struct {
	ID          string   `json:"id"`
	ExternalID  string   `json:"external_id"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *int     `json:"quantity"`
}

var db *sql.DB

func main() {
	fmt.Println("Connecting to database...")

	dsn := "postgresql://postgres:postgres@localhost:5432/mdg?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening connection to database", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	fmt.Println("Connection to dabase succesfully established!")

	fmt.Println("Starting server on port 8000")

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/items/", handleItems)
	http.HandleFunc("/api/item/", handleItem)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "status: ok")
	case "POST":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetItems(w, r)
	case "POST":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetItem(w, r)
	case "POST":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := getItems()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func getItems() ([]Item, error) {
	var items []Item
	rows, err := db.Query("SELECT * FROM item LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("allItems: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			item     Item
			price    sql.NullFloat64
			quantity sql.NullInt64
		)

		if err := rows.Scan(&item.ID, &item.ExternalID, &item.Description, &price, &quantity); err != nil {
			return nil, fmt.Errorf("allItems: %v", err)
		}

		if price.Valid {
			item.Price = &price.Float64
		}

		if quantity.Valid {
			qty := int(quantity.Int64)
			item.Quantity = &qty
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("allItems: %v", err)
	}

	return items, nil
}

func handleGetItem(w http.ResponseWriter, r *http.Request) {
	path, err := parsePathParam(r, "/api/item/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "You requested item ID: %s", path)
}

func parsePathParam(r *http.Request, routePrefix string) (string, error) {
	path := strings.TrimPrefix(r.URL.Path, routePrefix)

	if path == "" || path == "/" {
		return "", fmt.Errorf("parameter is required")
	}

	return path, nil
}
