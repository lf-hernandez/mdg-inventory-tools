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

type App struct {
	Router  *http.ServeMux
	DB      *sql.DB
	Cors    *cors.Cors
	Address string
}

func (app *App) Init(cfg *config.Config) error {
	utils.LogInfo("Connecting to %v", cfg.Database.URL)
	var err error
	app.DB, err = sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		utils.LogError(fmt.Errorf("error opening connection to database: %v", err))
		os.Exit(1)
	}

	for i := 0; i < cfg.Database.MaxRetry; i++ {
		err = app.DB.Ping()
		if err == nil {
			utils.LogInfo("Connection established!")
			break
		}
		utils.LogInfo("Attempt %d: Error connecting to database: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		utils.LogFatal("Failed to connect to database after %d attempts: %v", cfg.Database.MaxRetry, err)
		return err
	}

	port := cfg.Port
	if port == "" {
		utils.LogInfo("$PORT not set, defaulting to 8000")
		cfg.Port = "8000"
	}
	app.Address = ":" + cfg.Port

	app.Cors = cors.New(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	app.Router = initRouter(handlers.NewHandlerDependencies(app.DB, cfg.JWTSecret))

	return nil
}

func (app *App) Run() {
	defer app.DB.Close()
	utils.LogInfo("Starting server on port %s", app.Address)
	log.Fatal(http.ListenAndServe(app.Address, app.Cors.Handler(app.Router)))
}
