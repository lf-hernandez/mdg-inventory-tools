package main

import "time"

type Config struct {
	DatabaseURL string
	Port        string
	CORSOrigins []string
}

type Item struct {
	ID                string    `json:"id"`
	PartNumber        string    `json:"partNumber"`
	SerialNumber      string    `json:"serialNumber"`
	PurchaseOrder     string    `json:"purchaseOrder"`
	Description       string    `json:"description"`
	Category          string    `json:"category"`
	Price             *float64  `json:"price"`
	Quantity          *int      `json:"quantity"`
	Status            string    `json:"status"`
	RepairOrderNumber string    `json:"repair_order_number"`
	Condition         string    `json:"condition"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
}

type Inventory struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
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
