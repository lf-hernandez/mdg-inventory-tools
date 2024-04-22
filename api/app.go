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
	Config  *config.Config
	Address string
}

func (app *App) Init() error {
	app.Config = config.LoadConfig()

	utils.LogInfo("Connecting to %v", app.Config.Database.URL)
	var err error
	app.DB, err = sql.Open("postgres", app.Config.Database.URL)
	if err != nil {
		utils.LogError(fmt.Errorf("error opening connection to database: %v", err))
		os.Exit(1)
	}
	defer app.DB.Close()

	for i := 0; i < app.Config.Database.MaxRetry; i++ {
		err = app.DB.Ping()
		if err == nil {
			utils.LogInfo("Connection established!")
			break
		}
		utils.LogInfo("Attempt %d: Error connecting to database: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		utils.LogFatal("Failed to connect to database after %d attempts: %v", app.Config.Database.MaxRetry, err)
		return err
	}

	port := app.Config.Port
	if port == "" {
		utils.LogInfo("$PORT not set, defaulting to 8000")
		app.Config.Port = "8000"
	}
	app.Address = ":" + app.Config.Port

	app.Cors = cors.New(cors.Options{
		AllowedOrigins:   app.Config.CORSOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	app.Router = initRouter(handlers.NewHandlerDependencies(app.DB, app.Config.JWTSecret))

	return nil
}

func (app *App) Run() {
	utils.LogInfo("Starting server on port %s", app.Config.Port)
	log.Fatal(http.ListenAndServe(app.Address, app.Cors.Handler(app.Router)))
}
