package models

import "time"

type Item struct {
	ID                string   `json:"id"`
	PartNumber        string   `json:"partNumber"`
	SerialNumber      string   `json:"serialNumber"`
	PurchaseOrder     string   `json:"purchaseOrder"`
	Description       string   `json:"description"`
	Category          string   `json:"category"`
	Price             *float64 `json:"price"`
	Quantity          *int     `json:"quantity"`
	Status            string   `json:"status"`
	RepairOrderNumber string   `json:"repair_order_number"`
	Condition         string   `json:"condition"`
	Location          string   `json:"location"`
	Notes             string   `json:"notes"`

	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
