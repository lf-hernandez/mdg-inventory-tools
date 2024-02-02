package main

import (
	"net/http"

	"github.com/lf-hernandez/mdg-inventory-tools/api/handlers"
	"github.com/lf-hernandez/mdg-inventory-tools/api/middleware"
)

func initRouter(deps *handlers.HandlerDependencies) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/login", http.HandlerFunc(deps.HandleLogin))
	mux.Handle("/api/signup", http.HandlerFunc(deps.HandleSignup))

	// Pass the JWT secret to JwtMiddleware
	mux.Handle("/api/update-password", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(deps.HandleUpdatePassword)))
	mux.Handle("/api/items", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeItems(deps, w, r)
	})))
	mux.Handle("/api/items/", middleware.JwtMiddleware(deps.JwtSecret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeSpecificItem(deps, w, r)
	})))

	return mux
}
