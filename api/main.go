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
	utils.LogInfo(cfg.Database.URL)

	var err error
	db, err = sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		utils.LogError(fmt.Errorf("error opening connection to database: %v", err))
		os.Exit(1)
	}
	defer db.Close()

	for i := 0; i < cfg.Database.MaxRetry; i++ {
		err = db.Ping()
		if err == nil {
			utils.LogInfo("Connection to database successfully established!")
			break
		}
		utils.LogInfo("Attempt %d: Error connecting to database: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		utils.LogFatal("Failed to connect to database after %d attempts: %v", cfg.Database.MaxRetry, err)
	}

	utils.LogInfo("Attempting to bind to port")
	port := cfg.Port
	if port == "" {
		utils.LogInfo("PORT not set, defaulting to 8000")
		port = "8000"
	}

	utils.LogInfo("Starting server on port %v", port)
	utils.LogInfo("cors origins: %v ", cfg.CORSOrigins)

	router := initRouter(handlers.NewHandlerDependencies(db, cfg.JWTSecret))
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	utils.LogInfo("Starting server on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, corsHandler.Handler(router)))
}
