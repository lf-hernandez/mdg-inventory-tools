package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lf-hernandez/mdg-inventory-tools/api/auth"
	"github.com/lf-hernandez/mdg-inventory-tools/api/utils"
)

func JwtMiddleware(jwtSecret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := auth.ExtractJwtToken(r)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			utils.LogError(fmt.Errorf("error parsing token: %w", err))
			http.Error(w, "Unauthorized - token invalid", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			utils.LogError(fmt.Errorf("invalid token"))
			http.Error(w, "Unauthorized - token not valid", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
