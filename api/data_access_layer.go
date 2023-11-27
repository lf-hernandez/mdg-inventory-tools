package main

import (
	"database/sql"
	"fmt"
)

func fetchDbItems() ([]Item, error) {
	var items []Item
	rows, err := db.Query("SELECT * FROM item LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			item          Item
			price         sql.NullFloat64
			quantity      sql.NullInt64
			serialNumber  sql.NullString
			purchaseOrder sql.NullString
			category      sql.NullString
			inventoryID   sql.NullString
		)

		if err := rows.Scan(
			&item.ID,
			&item.PartNumber,
			&serialNumber,
			&category,
			&item.Description,
			&price,
			&quantity,
			&purchaseOrder,
			&inventoryID); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
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
		if inventoryID.Valid {
			item.InventoryID = inventoryID.String
		}
		if price.Valid {
			item.Price = &price.Float64
		}
		if quantity.Valid {
			qty := int(quantity.Int64)
			item.Quantity = &qty
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in fetch and iteration over rows: %w", err)
	}

	return items, nil
}

func fetchDbItem(partNumber string) (Item, error) {
	var (
		item          Item
		price         sql.NullFloat64
		quantity      sql.NullInt64
		serialNumber  sql.NullString
		purchaseOrder sql.NullString
		category      sql.NullString
		inventoryID   sql.NullString
	)
	err := db.
		QueryRow("SELECT * FROM item WHERE part_number = $1", partNumber).
		Scan(
			&item.ID,
			&item.PartNumber,
			&serialNumber,
			&category,
			&item.Description,
			&price,
			&quantity,
			&purchaseOrder,
			&inventoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, fmt.Errorf("no item found with part number %s: %w", partNumber, err)
		}
		return Item{}, fmt.Errorf("error executing query: %w", err)
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
	if inventoryID.Valid {
		item.InventoryID = inventoryID.String
	}
	if price.Valid {
		item.Price = &price.Float64
	}
	if quantity.Valid {
		qty := int(quantity.Int64)
		item.Quantity = &qty
	}

	return item, nil
}

func updateDbItem(item *Item) error {
	if err := validateItem(item); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	stmt, err := db.Prepare("UPDATE item SET description = $1, price = $2, quantity = $3, category = $4 WHERE id = $5")
	if err != nil {
		return fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Description, item.Price, item.Quantity, item.Category, item.ID)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}

func createDbItem(item Item) (Item, error) {
	if err := validateItem(&item); err != nil {
		return Item{}, fmt.Errorf("validation failed: %w", err)
	}

	stmt, err := db.Prepare("INSERT INTO item (part_number, serial_number, category, description, price, quantity, purchase_order, inventory_id) VALUES ($1, $2, $3, $4, $6, $7) RETURNING id")
	if err != nil {
		return Item{}, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	var id string
	err = stmt.
		QueryRow(
			item.PartNumber,
			item.SerialNumber,
			item.Category,
			item.Description,
			item.Price,
			item.Quantity,
			item.PurchaseOrder).
		Scan(&id)
	if err != nil {
		return Item{}, fmt.Errorf("error executing SQL statement: %w", err)
	}

	item.ID = id
	return item, nil
}
