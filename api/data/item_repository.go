package data

import (
	"database/sql"
	"fmt"

	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
	"github.com/lf-hernandez/mdg-inventory-tools/api/utils"
)

type ItemRepository struct {
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (repo *ItemRepository) FetchTotalItemCount() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM item`
	err := repo.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error fetching total item count: %w", err)
	}
	return count, nil
}

func (repo *ItemRepository) FetchDbItems(page int, limit int) ([]models.Item, error) {
	var items []models.Item

	var query string
	var args []interface{}

	if page < 0 && limit < 0 {
		query = `
			SELECT
				id,
				part_number,
				description,
				price,
				quantity,
				serial_number,
				purchase_order,
				category,
				status,
				repair_order_number,
				condition,
				location,
				notes
			FROM item
			ORDER BY id`
	} else {
		offset := (page - 1) * limit
		query = `
			SELECT
				id,
				part_number,
				description,
				price,
				quantity,
				serial_number,
				purchase_order,
				category,
				status,
				repair_order_number,
				condition,
				location,
				notes
			FROM item
			ORDER BY id
			LIMIT $1 OFFSET $2`
		args = append(args, limit, offset)
	}

	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			item              models.Item
			description       sql.NullString
			price             sql.NullFloat64
			quantity          sql.NullInt64
			serialNumber      sql.NullString
			purchaseOrder     sql.NullString
			category          sql.NullString
			status            sql.NullString
			repairOrderNumber sql.NullString
			condition         sql.NullString
			location          sql.NullString
			notes             sql.NullString
		)

		if err := rows.Scan(
			&item.ID,
			&item.PartNumber,
			&description,
			&price,
			&quantity,
			&serialNumber,
			&purchaseOrder,
			&category,
			&status,
			&repairOrderNumber,
			&condition,
			&location,
			&notes); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		if description.Valid {
			item.Description = description.String
		}
		if price.Valid {
			item.Price = &price.Float64
		}
		if quantity.Valid {
			qty := int(quantity.Int64)
			item.Quantity = &qty
		}
		if serialNumber.Valid {
			item.SerialNumber = serialNumber.String
		}
		if purchaseOrder.Valid {
			item.PurchaseOrder = purchaseOrder.String
		}
		if category.Valid {
			item.Category = category.String
		}
		if status.Valid {
			item.Status = status.String
		}
		if repairOrderNumber.Valid {
			item.RepairOrderNumber = repairOrderNumber.String
		}
		if condition.Valid {
			item.Condition = condition.String
		}
		if location.Valid {
			item.Location = location.String
		}
		if notes.Valid {
			item.Notes = notes.String
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in fetch and iteration over rows: %w", err)
	}

	return items, nil
}

func (repo *ItemRepository) FetchDbItemsWithSearch(searchQuery string) ([]models.Item, error) {
	var items []models.Item

	likeQuery := "%" + searchQuery + "%"

	query := `SELECT
				id,
				part_number,
				description,
				price,
				quantity,
				serial_number,
				purchase_order,
				category,
				status,
				repair_order_number,
				condition,
				location,
				notes
			FROM item
			WHERE part_number = $1
			OR serial_number = $1
			OR description ILIKE $2`
	rows, err := repo.DB.Query(query, searchQuery, likeQuery)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			item              models.Item
			description       sql.NullString
			price             sql.NullFloat64
			quantity          sql.NullInt64
			serialNumber      sql.NullString
			purchaseOrder     sql.NullString
			category          sql.NullString
			status            sql.NullString
			repairOrderNumber sql.NullString
			condition         sql.NullString
			location          sql.NullString
			notes             sql.NullString
		)

		if err := rows.Scan(
			&item.ID,
			&item.PartNumber,
			&description,
			&price,
			&quantity,
			&serialNumber,
			&purchaseOrder,
			&category,
			&status,
			&repairOrderNumber,
			&condition,
			&location,
			&notes); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		if description.Valid {
			item.Description = description.String
		}
		if price.Valid {
			item.Price = &price.Float64
		}
		if quantity.Valid {
			qty := int(quantity.Int64)
			item.Quantity = &qty
		}
		if serialNumber.Valid {
			item.SerialNumber = serialNumber.String
		}
		if purchaseOrder.Valid {
			item.PurchaseOrder = purchaseOrder.String
		}
		if category.Valid {
			item.Category = category.String
		}
		if status.Valid {
			item.Status = status.String
		}
		if repairOrderNumber.Valid {
			item.RepairOrderNumber = repairOrderNumber.String
		}
		if condition.Valid {
			item.Condition = condition.String
		}
		if location.Valid {
			item.Location = location.String
		}
		if notes.Valid {
			item.Notes = notes.String
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in fetch and iteration over rows: %w", err)
	}

	return items, nil
}

func (repo *ItemRepository) FetchDbItem(partNumber string) (models.Item, error) {
	var (
		item              models.Item
		description       sql.NullString
		price             sql.NullFloat64
		quantity          sql.NullInt64
		serialNumber      sql.NullString
		purchaseOrder     sql.NullString
		category          sql.NullString
		status            sql.NullString
		repairOrderNumber sql.NullString
		condition         sql.NullString
		location          sql.NullString
		notes             sql.NullString
	)
	err := repo.DB.
		QueryRow(`
		SELECT
			id,
			part_number,
			description,
			price,
			quantity,
			serial_number,
			purchase_order,
			category,
			status,
			repair_order_number,
			condition,
			location,
			notes
		FROM item
		WHERE part_number = $1`, partNumber).
		Scan(
			&item.ID,
			&item.PartNumber,
			&description,
			&price,
			&quantity,
			&serialNumber,
			&purchaseOrder,
			&category,
			&status,
			&repairOrderNumber,
			&condition,
			&location,
			&notes)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Item{}, fmt.Errorf("no item found with part number %s: %w", partNumber, err)
		}
		return models.Item{}, fmt.Errorf("error executing query: %w", err)
	}

	if description.Valid {
		item.Description = description.String
	}
	if price.Valid {
		item.Price = &price.Float64
	}
	if quantity.Valid {
		qty := int(quantity.Int64)
		item.Quantity = &qty
	}
	if serialNumber.Valid {
		item.SerialNumber = serialNumber.String
	}
	if purchaseOrder.Valid {
		item.PurchaseOrder = purchaseOrder.String
	}
	if category.Valid {
		item.Category = category.String
	}
	if status.Valid {
		item.Status = status.String
	}
	if repairOrderNumber.Valid {
		item.RepairOrderNumber = repairOrderNumber.String
	}
	if condition.Valid {
		item.Condition = condition.String
	}
	if location.Valid {
		item.Location = location.String
	}
	if notes.Valid {
		item.Notes = notes.String
	}

	return item, nil
}

func (repo *ItemRepository) UpdateDbItem(item *models.Item) error {
	if err := utils.ValidateItem(item); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	stmt, err := repo.DB.Prepare(`
		UPDATE item
		SET description = $1,
			price = $2,
			quantity = $3,
			serial_number = $4,
			purchase_order = $5,
			category = $6,
			status = $7,
			repair_order_number = $8,
			condition = $9,
			location = $10,
			notes = $11
		WHERE id = $12`)
	if err != nil {
		return fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		item.Description,
		item.Price,
		item.Quantity,
		item.SerialNumber,
		item.PurchaseOrder,
		item.Category,
		item.Status,
		item.RepairOrderNumber,
		item.Condition,
		item.Location,
		item.Notes,
		item.ID)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}

func (repo *ItemRepository) CreateDbItem(item models.Item) (models.Item, error) {
	if err := utils.ValidateItem(&item); err != nil {
		return models.Item{}, fmt.Errorf("validation failed: %w", err)
	}

	stmt, err := repo.DB.Prepare(`
		INSERT INTO item (
			part_number,
			description,
			price,
			quantity,
			serial_number,
			purchase_order,
			category,
			status,
			repair_order_number,
			condition,
			location,
			notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`)

	if err != nil {
		return models.Item{}, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	var id string
	err = stmt.
		QueryRow(
			item.PartNumber,
			item.Description,
			item.Price,
			item.Quantity,
			item.SerialNumber,
			item.PurchaseOrder,
			item.Category,
			item.Status,
			item.RepairOrderNumber,
			item.Condition,
			item.Location,
			item.Notes).
		Scan(&id)
	if err != nil {
		return models.Item{}, fmt.Errorf("error executing SQL statement: %w", err)
	}

	item.ID = id
	return item, nil
}
