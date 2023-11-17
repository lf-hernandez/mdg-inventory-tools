package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	items, err := allItems()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func allItems() ([]Item, error) {
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
