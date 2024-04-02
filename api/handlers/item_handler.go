package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/lf-hernandez/mdg-inventory-tools/api/data"
	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
	"github.com/lf-hernandez/mdg-inventory-tools/api/utils"
)

func (deps *HandlerDependencies) HandleGetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repo := data.NewItemRepository(deps.DB)

	searchQuery := r.URL.Query().Get("search")
	if searchQuery != "" {
		items, err := repo.FetchDbItemsWithSearch(searchQuery)
		if err != nil {
			utils.LogError(fmt.Errorf("error fetching items with search query '%s': %w", searchQuery, err))
			http.Error(w, fmt.Sprintf("Error fetching items: %v", err), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	page, limit := 1, 10

	if p := r.URL.Query().Get("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil {
			limit = parsedLimit
		}
	}

	totalCount, err := repo.FetchTotalItemCount()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching total item count: %v", err), http.StatusInternalServerError)
		return
	}

	items, err := repo.FetchDbItems(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := GetItemsResponse{
		Items:      items,
		TotalCount: totalCount,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (deps *HandlerDependencies) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId := r.PathValue("id")

	repo := data.NewItemRepository(deps.DB)
	item, err := repo.FetchDbItem(itemId)
	if err != nil {
		var statusCode int
		switch err {
		case sql.ErrNoRows:
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, fmt.Sprintf("Error fetching item: %v", err), statusCode)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (deps *HandlerDependencies) HandleUpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repo := data.NewItemRepository(deps.DB)

	itemId := r.PathValue("id")

	var updatedItem models.Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedItem.ID = itemId
	err = repo.UpdateDbItem(&updatedItem)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error updating item: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
}

func (deps *HandlerDependencies) HandleCreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repo := data.NewItemRepository(deps.DB)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newItem models.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		utils.LogError(fmt.Errorf("error creating new item: %w", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.ValidateItem(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdItem, err := repo.CreateDbItem(newItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating item: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdItem)
}
