package main

import (
	"net/http"
)

func initRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/login", http.HandlerFunc(handleLogin))
	mux.Handle("/api/signup", http.HandlerFunc(handleSignup))
	mux.Handle("/api/items", JwtMiddleware(http.HandlerFunc(routeItems)))
	mux.Handle("/api/items/", JwtMiddleware(http.HandlerFunc(routeSpecificItem)))
	return mux
}
