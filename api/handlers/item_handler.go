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
	repo := data.NewItemRepository(deps.DB)

	searchQuery := r.URL.Query().Get("search")
	if searchQuery != "" {
		items, err := repo.FetchDbItemsWithSearch(searchQuery)
		if err != nil {
			utils.LogError(fmt.Errorf("error fetching items with search query '%s': %w", searchQuery, err))
			http.Error(w, "Error getting items", http.StatusInternalServerError)
			return
		}

		err = utils.WriteJSONResponse(w, http.StatusOK, items, nil)
		if err != nil {
			utils.LogError(fmt.Errorf("get items json response error: %w", err))
			http.Error(w, "Error getting items", http.StatusInternalServerError)
			return
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
		utils.LogError(fmt.Errorf("fetch total item count error: %w", err))
		http.Error(w, "Error getting items", http.StatusInternalServerError)
		return
	}

	items, err := repo.FetchDbItems(page, limit)
	if err != nil {
		utils.LogError(fmt.Errorf("get items error: %w", err))
		http.Error(w, "Error getting items", http.StatusInternalServerError)
		return
	}

	response := GetItemsResponse{
		Items:      items,
		TotalCount: totalCount,
	}

	err = utils.WriteJSONResponse(w, http.StatusOK, response, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("get items json response error: %w", err))
		http.Error(w, "Error getting items", http.StatusInternalServerError)
		return
	}
}

func (deps *HandlerDependencies) HandleGetItem(w http.ResponseWriter, r *http.Request) {
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
		utils.LogError(fmt.Errorf("fetch item error: %w", err))
		http.Error(w, "Error getting item", statusCode)
		return
	}

	err = utils.WriteJSONResponse(w, http.StatusOK, item, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("get item json response error: %w", err))
		http.Error(w, "Error getting item", http.StatusInternalServerError)
		return
	}
}

func (deps *HandlerDependencies) HandleUpdateItem(w http.ResponseWriter, r *http.Request) {
	repo := data.NewItemRepository(deps.DB)

	var updatedItem models.Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		utils.LogError(fmt.Errorf("decode update item error: %w", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	itemId := r.PathValue("id")
	updatedItem.ID = itemId
	err = repo.UpdateDbItem(&updatedItem)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.LogError(fmt.Errorf("update item not found error: %w", err))
			http.Error(w, "Error updating item", http.StatusNotFound)
			return
		} else {
			utils.LogError(fmt.Errorf("update item repo error: %w", err))
			http.Error(w, fmt.Sprintf("Error updating item: %v", err), http.StatusInternalServerError)
			return
		}
	}

	err = utils.WriteJSONResponse(w, http.StatusOK, updatedItem, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("update item json response error: %w", err))
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}
}

func (deps *HandlerDependencies) HandleCreateItem(w http.ResponseWriter, r *http.Request) {
	repo := data.NewItemRepository(deps.DB)

	var newItem models.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		utils.LogError(fmt.Errorf("decode new item error: %w", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := utils.ValidateItem(&newItem); err != nil {
		utils.LogError(fmt.Errorf("validate new item error: %w", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	createdItem, err := repo.CreateDbItem(newItem)
	if err != nil {
		utils.LogError(fmt.Errorf("create item error: %w", err))
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSONResponse(w, http.StatusCreated, createdItem, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("create item json response error: %w", err))
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}
}
