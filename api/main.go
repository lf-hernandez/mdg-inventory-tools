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
	case http.MethodGet:
		fmt.Fprintf(w, "status: ok")
	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func routeItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetItems(w, r)
	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func routeItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetItem(w, r)
	case http.MethodPost:
		handleCreateItem(w, r)
	case http.MethodPut:
		handleUpdateItem(w, r)
	case http.MethodDelete:
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

func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := extractPathParam(r, "/api/item")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedItem Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedItem.ID = itemId
	err = updateDbItem(&updatedItem)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error updating item: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
}

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateItem(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdItem, err := createDbItem(newItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating item: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdItem)
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

func updateDbItem(item *Item) error {
	if err := validateItem(item); err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE item SET description = $1, price = $2, quantity = $3 WHERE id = $4")
	if err != nil {
		return fmt.Errorf("updateDbItem: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Description, item.Price, item.Quantity, item.ID)
	if err != nil {
		return fmt.Errorf("e rror updating item: %v", err)
	}

	return nil
}

func createDbItem(item Item) (Item, error) {
	stmt, err := db.Prepare("INSERT INTO item (external_id, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return Item{}, fmt.Errorf("createDbItem: %v", err)
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(item.ExternalID, item.Description, item.Price, item.Quantity).Scan(&id)
	if err != nil {
		return Item{}, fmt.Errorf("createDbItem: %v", err)
	}

	item.ID = id
	return item, nil
}

func extractPathParam(r *http.Request, routePrefix string) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, routePrefix)

	if param == "" || param == "/" {
		return "", fmt.Errorf("parameter is required")
	}

	return param, nil
}

func validateItem(item *Item) error {
	if item.ID == "" {
		return fmt.Errorf("ID is required")
	}

	if item.ExternalID == "" {
		return fmt.Errorf("external ID is required")
	}

	if item.Price == nil {
		return fmt.Errorf("price is requried")
	}

	if item.Price != nil && *item.Price < 0 {
		return fmt.Errorf("price must be non-negative")
	}

	if item.Quantity == nil {
		return fmt.Errorf("quantity is required")
	}

	return nil
}
