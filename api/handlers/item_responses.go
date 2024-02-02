package handlers

import "github.com/lf-hernandez/mdg-inventory-tools/api/models"

type GetItemsResponse struct {
	Items      []models.Item `json:"items"`
	TotalCount int           `json:"totalCount"`
}
