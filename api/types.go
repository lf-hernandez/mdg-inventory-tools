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
}
