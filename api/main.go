package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	fmt.Println("Connecting to database.")

	dsn := os.Getenv("POSTGRES_DSN")
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

	port := os.Getenv("PORT")
	if port == "" {
		logInfo("PORT not set, defaulting to 8000")
		port = "8000"
	}

	fmt.Printf("Starting server on port %v", port)

	http.HandleFunc("/api/login", handleLogin)
	http.HandleFunc("/api/signup", handleSignup)
	http.Handle("/api/items", JwtMiddleware(http.HandlerFunc(routeItems)))
	http.Handle("/api/items/", JwtMiddleware(http.HandlerFunc(routeSpecificItem)))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
