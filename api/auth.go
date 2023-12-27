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

func authenticateUser(r *http.Request) (string, error) {
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
		return "", fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error extracting claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("user_id not found in token")
	}

	return userID, nil
}
