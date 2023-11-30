package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJwtSecret() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT secret not set")
	}
	return jwtSecret, nil
}

func createToken(userId string) (string, error) {
	jwtSecret, err := getJwtSecret()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func extractJwtToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	return strings.TrimSpace(strings.Replace(bearerToken, "Bearer ", "", 1))
}
