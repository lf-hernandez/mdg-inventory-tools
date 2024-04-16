package handlers

import "github.com/lf-hernandez/mdg-inventory-tools/api/models"

type UserResponse struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  models.Role `json:"role"`
}
