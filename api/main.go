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

	http.HandleFunc("/", routeRoot)
	http.HandleFunc("/api/items/", routeItems)
	http.HandleFunc("/api/item/", routeItem)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func routeRoot(w http.ResponseWriter, r *http.Request) {
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

func routeItems(w http.ResponseWriter, r *http.Request) {
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

func routeItem(w http.ResponseWriter, r *http.Request) {
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

	items, err := fetchDbItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := extractPathParam(r, "/api/item/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := fetchDbItem(itemId)
	if err != nil {
		var statusCode int
		switch err {
		case sql.ErrNoRows:
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, fmt.Sprintf("Error fetching item: %v", err), statusCode)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchDbItems() ([]Item, error) {
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

func fetchDbItem(itemId string) (Item, error) {
	var item Item
	var price sql.NullFloat64
	var quantity sql.NullInt64

	err := db.QueryRow("SELECT * FROM item WHERE external_id = $1", itemId).Scan(&item.ID, &item.ExternalID, &item.Description, &price, &quantity)
	if err != nil {
		return Item{}, err
	}

	if price.Valid {
		item.Price = &price.Float64
	}

	if quantity.Valid {
		qty := int(quantity.Int64)
		item.Quantity = &qty
	}

	return item, nil
}

func extractPathParam(r *http.Request, routePrefix string) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, routePrefix)

	if param == "" || param == "/" {
		return "", fmt.Errorf("parameter is required")
	}

	return param, nil
}
