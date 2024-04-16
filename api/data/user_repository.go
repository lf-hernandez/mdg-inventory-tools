package data

import (
	"database/sql"
	"fmt"

	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
	"github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user models.User, inventories []models.Inventory) (models.User, error) {
	inventoryNames := make([]string, len(inventories))
	for i, inventory := range inventories {
		inventoryNames[i] = inventory.Name
	}
	stmt, err := repo.DB.Prepare(`
	INSERT INTO app_user (
		name,
		email,
		password,
		role,
		inventories
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`)

	if err != nil {
		return models.User{}, fmt.Errorf("error preparing statement: %w", err)
	}

	defer stmt.Close()

	var id string

	err = stmt.QueryRow(user.Name, user.Email, user.Password, user.Role, pq.Array(inventoryNames)).Scan(&id)
	if err != nil {
		return models.User{}, fmt.Errorf("error executing SQL statement: %w", err)
	}

	user.ID = id
	return user, nil
}

func (repo *UserRepository) FetchUserByEmail(email string) (models.User, error) {
	var user models.User

	err := repo.DB.QueryRow(
		"SELECT id, name, email, password, role FROM app_user WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with email %s: %w", email, err)
		}
		return models.User{}, fmt.Errorf("error executing query: %w", err)
	}

	return user, nil
}

func (repo *UserRepository) FetchUserByID(userID string) (models.User, error) {
	var user models.User

	err := repo.DB.QueryRow(
		"SELECT id, name, email, password FROM app_user WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with ID %s: %w", userID, err)
		}
		return models.User{}, fmt.Errorf("error executing query: %w", err)
	}

	return user, nil
}

func (repo *UserRepository) UpdatePassword(userID, newPassword string) error {
	stmt, err := repo.DB.Prepare("UPDATE app_user SET password = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(newPassword, userID)
	if err != nil {
		return fmt.Errorf("error executing update statement: %w", err)
	}

	return nil
}
