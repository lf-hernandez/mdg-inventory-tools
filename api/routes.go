package main

import (
	"net/http"

	"github.com/lf-hernandez/mdg-inventory-tools/api/handlers"
	"github.com/lf-hernandez/mdg-inventory-tools/api/middleware"
)

func initRouter(deps *handlers.HandlerDependencies) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /api/auth/login", http.HandlerFunc(deps.HandleLogin))
	// mux.Handle("POST /api/auth/signup", http.HandlerFunc(deps.HandleSignup))

	mux.Handle("PUT /api/account/update-password", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleUpdatePassword)))

	mux.Handle("GET /api/items", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleGetItems)))
	mux.Handle("POST /api/items", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleCreateItem)))
	mux.Handle("GET /api/items/{id}", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleGetItem)))
	mux.Handle("PUT /api/items/{id}", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleUpdateItem)))
	mux.Handle("GET /api/items/export", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleExportCSV)))

	return mux
}
