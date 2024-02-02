package handlers

import (
	"database/sql"
)

type HandlerDependencies struct {
	DB        *sql.DB
	JwtSecret string
}

func NewHandlerDependencies(db *sql.DB, jwtSecret string) *HandlerDependencies {
	return &HandlerDependencies{
		DB:        db,
		JwtSecret: jwtSecret,
	}
}
