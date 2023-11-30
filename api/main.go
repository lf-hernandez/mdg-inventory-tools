package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	fmt.Println("Connecting to database.")

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

	fmt.Println("Connection to database successfully established!")

	fmt.Println("Starting server on port 8000")

	http.HandleFunc("/api/login", handleLogin)
	http.HandleFunc("/api/signup", handleSignup)
	http.Handle("/api/items", JwtMiddleware(http.HandlerFunc(routeItems)))
	http.Handle("/api/items/", JwtMiddleware(http.HandlerFunc(routeSpecificItem)))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
