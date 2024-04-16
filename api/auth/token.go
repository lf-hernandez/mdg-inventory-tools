package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
)

func CreateToken(userId string, role models.Role, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractJwtToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	return strings.TrimSpace(strings.Replace(bearerToken, "Bearer ", "", 1))
}

func AuthenticateUser(r *http.Request, jwtSecret string) (string, models.Role, error) {
	tokenString := ExtractJwtToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("error extracting claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", "", fmt.Errorf("user_id not found in token")
	}

	role, ok := claims["role"].(models.Role)
	if !ok {
		return "", "", fmt.Errorf("role not found in token")
	}

	return userID, role, nil
}
