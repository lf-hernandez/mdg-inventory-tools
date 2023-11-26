package main

import (
	"database/sql"
	"fmt"
)

func fetchDbItems() ([]Item, error) {
	var items []Item
	rows, err := db.Query("SELECT * FROM item LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("error exectuting query: %v", err)
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
			return nil, fmt.Errorf("error scanning row: %v", err)
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
		return nil, fmt.Errorf("error in fetch and iteration over rows: %v", err)
	}

	return items, nil
}

func fetchDbItem(itemId string) (Item, error) {
	var item Item
	var price sql.NullFloat64
	var quantity sql.NullInt64

	err := db.
		QueryRow("SELECT * FROM item WHERE part_number = $1", itemId).
		Scan(
			&item.ID,
			&item.PartNumber,
			&item.SerialNumber,
			&item.Category,
			&item.Description,
			&price,
			&quantity,
			&item.PurchaseOrder)
	if err != nil {
		return Item{}, err
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
		return err
	}

	stmt, err := db.Prepare("UPDATE item SET description = $1, price = $2, quantity = $3 WHERE id = $4")
	if err != nil {
		return fmt.Errorf("updateDbItem: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Description, item.Price, item.Quantity, item.ID)
	if err != nil {
		return fmt.Errorf("e rror updating item: %v", err)
	}

	return nil
}

func createDbItem(item Item) (Item, error) {
	stmt, err := db.Prepare("INSERT INTO item (part_number, serial_number, category, description, price, quantity, purchase_order) VALUES ($1, $2, $3, $4, $6, $7) RETURNING id")
	if err != nil {
		return Item{}, fmt.Errorf("createDbItem: %v", err)
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
		return Item{}, fmt.Errorf("createDbItem: %v", err)
	}

	item.ID = id
	return item, nil
}
