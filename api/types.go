package main

type Item struct {
	ID            string   `json:"id"`
	PartNumber    string   `json:"part_number"`
	SerialNumber  string   `json:"serial_number"`
	PurchaseOrder string   `json:"purchase_order"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	Price         *float64 `json:"price"`
	Quantity      *int     `json:"quantity"`
	InventoryID   string   `json:"inventory_id"`
}

type Inventory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
