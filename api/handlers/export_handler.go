package handlers

import (
	"bytes"
	"encoding/csv"
	"net/http"
	"time"

	"github.com/lf-hernandez/mdg-inventory-tools/api/data"
	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
	"github.com/lf-hernandez/mdg-inventory-tools/api/utils"
)

func (deps *HandlerDependencies) HandleExportCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	filename := "mdg_inventory_" + time.Now().Format("2006-01-02") + ".csv"
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	repo := data.NewItemRepository(deps.DB)

	searchQuery := r.URL.Query().Get("search")
	var items []models.Item
	var err error
	if searchQuery != "" {
		items, err = repo.FetchDbItemsWithSearch(searchQuery)
	} else {
		items, err = repo.FetchDbItems(-1, -1)
	}

	if err != nil {
		http.Error(w, "Error fetching items: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var csvData bytes.Buffer
	writer := csv.NewWriter(&csvData)

	csvHeader := []string{
		"PartNumber",
		"SerialNumber",
		"PurchaseOrder",
		"Description",
		"Category",
		"Price",
		"Quantity",
		"Status",
		"RepairOrderNumber",
		"Condition",
	}
	writer.Write(csvHeader)

	for _, item := range items {
		writer.Write([]string{
			item.PartNumber,
			item.SerialNumber,
			item.PurchaseOrder,
			item.Description,
			item.Category,
			utils.FormatFloat(item.Price),
			utils.FormatInt(item.Quantity),
			item.Status,
			item.RepairOrderNumber,
			item.Condition,
		})
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		http.Error(w, "Error writing CSV: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(csvData.Bytes())
	if err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
