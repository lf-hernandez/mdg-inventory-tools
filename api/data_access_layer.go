package main

import (
	"database/sql"
	"fmt"
)

func fetchDbItems(page int, limit int) ([]Item, error) {
	var items []Item

	offset := (page - 1) * limit
	query := `
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
			condition
		FROM item
		ORDER BY id
		LIMIT $1 OFFSET $2`
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			item              Item
			description       sql.NullString
			price             sql.NullFloat64
			quantity          sql.NullInt64
			serialNumber      sql.NullString
			purchaseOrder     sql.NullString
			category          sql.NullString
			status            sql.NullString
			repairOrderNumber sql.NullString
			condition         sql.NullString
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
			&condition); err != nil {
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

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in fetch and iteration over rows: %w", err)
	}

	return items, nil
}

func fetchDbItemsWithSearch(searchQuery string) ([]Item, error) {
	var items []Item

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
				condition
			FROM item
			WHERE part_number = $1
			OR serial_number = $1
			OR description ILIKE $2`
	rows, err := db.Query(query, searchQuery, likeQuery)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			item              Item
			description       sql.NullString
			price             sql.NullFloat64
			quantity          sql.NullInt64
			serialNumber      sql.NullString
			purchaseOrder     sql.NullString
			category          sql.NullString
			status            sql.NullString
			repairOrderNumber sql.NullString
			condition         sql.NullString
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
			&condition); err != nil {
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

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in fetch and iteration over rows: %w", err)
	}

	return items, nil
}

func fetchDbItem(partNumber string) (Item, error) {
	var (
		item              Item
		description       sql.NullString
		price             sql.NullFloat64
		quantity          sql.NullInt64
		serialNumber      sql.NullString
		purchaseOrder     sql.NullString
		category          sql.NullString
		status            sql.NullString
		repairOrderNumber sql.NullString
		condition         sql.NullString
	)
	err := db.
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
			condition
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
			&condition)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, fmt.Errorf("no item found with part number %s: %w", partNumber, err)
		}
		return Item{}, fmt.Errorf("error executing query: %w", err)
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

	return item, nil
}

func updateDbItem(item *Item) error {
	if err := validateItem(item); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	stmt, err := db.Prepare(`
		UPDATE item
		SET description = $1,
			price = $2,
			quantity = $3,
			serial_number = $4,
			purchase_order = $5,
			category = $6,
			status = $7,
			repair_order_number = $8,
			condition = $9
		WHERE id = $10`)
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
		item.ID)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}

func createDbItem(item Item) (Item, error) {
	if err := validateItem(&item); err != nil {
		return Item{}, fmt.Errorf("validation failed: %w", err)
	}

	stmt, err := db.Prepare(`
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
			condition
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`)

	if err != nil {
		return Item{}, fmt.Errorf("error preparing statement: %w", err)
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
			item.Condition).
		Scan(&id)
	if err != nil {
		return Item{}, fmt.Errorf("error executing SQL statement: %w", err)
	}

	item.ID = id
	return item, nil
}

func createUser(user User) (User, error) {
	stmt, err := db.Prepare(`
	INSERT INTO app_user (
		name,
		email,
		password
	) VALUES ($1, $2, $3)
	RETURNING id
	`)

	if err != nil {
		return User{}, fmt.Errorf("error preparing statement: %w", err)
	}

	defer stmt.Close()

	var id string

	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return User{}, fmt.Errorf("error executing SQL statement: %w", err)
	}

	user.ID = id
	return user, nil
}

func fetchUserByEmail(email string) (User, error) {
	var user User

	err := db.QueryRow(
		"SELECT id, name, email, password FROM app_user WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("no user found with email %s: %w", email, err)
		}
		return User{}, fmt.Errorf("error executing query: %w", err)
	}

	return user, nil
}
