package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lf-hernandez/mdg-inventory-tools/api/config"
	"github.com/lf-hernandez/mdg-inventory-tools/api/handlers"
	"github.com/lf-hernandez/mdg-inventory-tools/api/utils"
	"github.com/rs/cors"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	cfg := config.LoadConfig()
	utils.LogInfo("Connecting to database")
	dsn := cfg.DatabaseURL
	utils.LogInfo(dsn)
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening connection to database: %v\nDSN: %s", err, dsn)
	}

	defer db.Close()

	for connectionAttempt := 0; connectionAttempt < 5; connectionAttempt++ {
		err = db.Ping()
		if err == nil {
			utils.LogInfo("Connection to database successfully established!")
			break
		}

		errorMessage := fmt.Errorf("attempt %d: Error connecting to database: %v", connectionAttempt+1, err)
		utils.LogError(errorMessage)
		time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
	}

	if err != nil {
		log.Fatalf("After several attempts, failed to connect to database: %v\n", err)
	}

	utils.LogInfo("Attempting to bind to port")
	port := cfg.Port
	if port == "" {
		utils.LogInfo("PORT not set, defaulting to 8000")
		port = "8000"
	}

	utils.LogInfo("Starting server on port %v", port)
	utils.LogInfo("cors origins: %v ", cfg.CORSOrigins)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT secret is not set")
	}

	deps := handlers.NewHandlerDependencies(db, jwtSecret)
	router := initRouter(deps)

	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(router)))
}
