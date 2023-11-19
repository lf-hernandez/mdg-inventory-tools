package main

type Item struct {
	ID          string   `json:"id"`
	ExternalID  string   `json:"external_id"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *int     `json:"quantity"`
}
