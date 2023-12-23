package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	cfg := loadConfig()
	logInfo("Connecting to database")
	dsn := cfg.DatabaseURL
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
	port := cfg.Port
	if port == "" {
		logInfo("PORT not set, defaulting to 8000")
		port = "8000"
	}

	logInfo("Starting server on port %v", port)
	logInfo("cors origins: %v ", cfg.CORSOrigins)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	router := initRouter()
	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(router)))
}
