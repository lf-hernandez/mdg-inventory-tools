package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_ORIGIN")},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	mux := http.NewServeMux()

	mux.Handle("/api/login", c.Handler(http.HandlerFunc(handleLogin)))
	mux.Handle("/api/signup", c.Handler(http.HandlerFunc(handleSignup)))
	mux.Handle("/api/items", c.Handler(JwtMiddleware(http.HandlerFunc(routeItems))))
	mux.Handle("/api/items/", c.Handler(JwtMiddleware(http.HandlerFunc(routeSpecificItem))))

	log.Fatal(http.ListenAndServe(":"+port, c.Handler(mux)))
}
