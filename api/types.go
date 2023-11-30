package main

import "time"

type Item struct {
	ID            string   `json:"id"`
	PartNumber    string   `json:"partNumber"`
	SerialNumber  string   `json:"serialNumber"`
	PurchaseOrder string   `json:"purchaseOrder"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	Price         *float64 `json:"price"`
	Quantity      *int     `json:"quantity"`
	InventoryID   string   `json:"inventoryId"`
}

type Inventory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
