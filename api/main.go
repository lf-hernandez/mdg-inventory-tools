package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	logInfo("Connecting to database")
	dsn := os.Getenv("DATABASE_URL")
	logInfo(dsn)
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening connection to database: %v\nDSN: %s", err, dsn)
	}

	defer db.Close()

	for connectionAttempt := 0; connectionAttempt < 5; connectionAttempt++ {
		err = db.Ping()
		if err == nil {
			logInfo("Connection to database successfully established!")
			break
		}

		errorMessage := fmt.Errorf("attempt %d: Error connecting to database: %v", connectionAttempt+1, err)
		logError(errorMessage)
		time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
	}

	if err != nil {
		log.Fatalf("After several attempts, failed to connect to database: %v\n", err)
	}

	logInfo("Attempting to bind to port")
	port := os.Getenv("PORT")
	if port == "" {
		logInfo("PORT not set, defaulting to 8000")
		port = "8000"
	}

	logInfo("Starting server on port %v", port)

	corsOrigins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")
	logInfo("cors origins: %v ", corsOrigins)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   corsOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	requestMultiplexer := http.NewServeMux()
	requestMultiplexer.Handle("/api/login", corsHandler.Handler(http.HandlerFunc(handleLogin)))
	requestMultiplexer.Handle("/api/signup", corsHandler.Handler(http.HandlerFunc(handleSignup)))
	requestMultiplexer.Handle("/api/items", corsHandler.Handler(JwtMiddleware(http.HandlerFunc(routeItems))))
	requestMultiplexer.Handle("/api/items/", corsHandler.Handler(JwtMiddleware(http.HandlerFunc(routeSpecificItem))))

	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(requestMultiplexer)))
}
