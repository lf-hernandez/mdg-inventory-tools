package main

import (
	"net/http"

	"github.com/lf-hernandez/mdg-inventory-tools/api/handlers"
	"github.com/lf-hernandez/mdg-inventory-tools/api/middleware"
)

func initRouter(deps *handlers.HandlerDependencies) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /api/login", http.HandlerFunc(deps.HandleLogin))
	mux.Handle("POST /api/signup", http.HandlerFunc(deps.HandleSignup))

	mux.Handle("POST /api/update-password", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleUpdatePassword)))

	mux.Handle("GET /api/items", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleGetItems)))
	mux.Handle("POST /api/items", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleCreateItem)))
	mux.Handle("GET /api/items/{id}", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleGetItem)))
	mux.Handle("PUT /api/items/{id}", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleUpdateItem)))
	mux.Handle("GET /api/items/csv", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleExportCSV)))

	return mux
}
