package handlers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
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
		utils.LogError(fmt.Errorf("export items records not found error: %w", err))
		http.Error(w, "Error exporting data", http.StatusInternalServerError)
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
			strconv.FormatFloat(*item.Price, 'f', 2, 64),
			strconv.Itoa(*item.Quantity),
			item.Status,
			item.RepairOrderNumber,
			item.Condition,
		})
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		utils.LogError(fmt.Errorf("export items error writing CSV: %w", err))
		http.Error(w, "Error exporting data", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(csvData.Bytes())
	if err != nil {
		utils.LogError(fmt.Errorf("export items json response error: %w", err))
		http.Error(w, "Error exporting data", http.StatusInternalServerError)
		return
	}
}
