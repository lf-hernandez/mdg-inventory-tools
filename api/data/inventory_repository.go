package data

import (
	"database/sql"
	"fmt"

	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
)

type InventoryRepository struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository {
	return &InventoryRepository{DB: db}
}

func (repo *InventoryRepository) Create(inventory models.Inventory) (models.Inventory, error) {
	stmt, err := repo.DB.Prepare(`
		INSERT INTO inventory (name)
		VALUES ($1)
		RETURNING id
	`)

	if err != nil {
		return models.Inventory{}, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	if err = stmt.QueryRow(inventory.Name).Scan(&inventory.ID, &inventory.CreatedAt, &inventory.ModifiedAt); err != nil {
		return models.Inventory{}, fmt.Errorf("error executing SQL statement: %w", err)
	}

	return inventory, nil
}

func (repo *InventoryRepository) List() ([]models.Inventory, error) {
	var inventories []models.Inventory

	rows, err := repo.DB.Query(`SELECT * FROM inventory`)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var inventory models.Inventory
		if err := rows.Scan(&inventory.ID, &inventory.Name, &inventory.CreatedAt, &inventory.ModifiedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting inventories: %w", err)
	}

	return inventories, nil
}

func (repo *InventoryRepository) Read(name string) (models.Inventory, error) {
	var inventory models.Inventory

	if err := repo.DB.QueryRow(`
		SELECT id, name, created_at, modified_at
		FROM inventory
		WHERE name = $1`, name).Scan(&inventory.ID, &inventory.Name, &inventory.CreatedAt, &inventory.ModifiedAt); err != nil {
		return models.Inventory{}, fmt.Errorf("error querying inventory: %w", err)
	}

	return inventory, nil
}

func (repo *InventoryRepository) Update(inventory *models.Inventory) error {
	stmt, err := repo.DB.Prepare(`
		UPDATE inventory
		SET name = $1
		WHERE id = $2
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(inventory.Name, inventory.ID)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}
