package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lf-hernandez/mdg-inventory-tools/api/auth"
	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.LogError(fmt.Errorf("unexpected token claims type"))
			http.Error(w, "Unauthorized - invalid token claims", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			utils.LogError(fmt.Errorf("role claim not found or not a string"))
			http.Error(w, "Unauthorized - role claim missing or invalid", http.StatusUnauthorized)
			return
		}
		resource, err := utils.ExtractResourceFromURL(r.URL.Path)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if !hasAccess(models.Role(role), r.Method, resource) {
			http.Error(w, "Forbidden - insufficient permissions", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func hasAccess(userRole models.Role, method string, resource string) bool {
	rolePermissions, ok := models.RoleResourcePermissions[userRole]
	if !ok {
		return false
	}

	resourcePermissions, ok := rolePermissions[resource]
	if !ok {
		return false
	}

	switch method {
	case http.MethodGet:
		return resourcePermissions["read"]
	case http.MethodPost:
		return resourcePermissions["create"]
	case http.MethodPut:
		return resourcePermissions["update"]
	case http.MethodDelete:
		return resourcePermissions["delete"]
	default:
		return false
	}
}
