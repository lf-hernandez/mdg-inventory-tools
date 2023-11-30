package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractJwtToken(r)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			jwtSecret, err := getJwtSecret()
			if err != nil {
				return nil, fmt.Errorf("secret not set")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			logError(fmt.Errorf("error parsing token: %w", err))
			http.Error(w, "Unauthorized - token invalid", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			logError(fmt.Errorf("invalid token"))
			http.Error(w, "Unauthorized - token not valid", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
