package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	searchQuery := r.URL.Query().Get("search")
	if searchQuery != "" {
		items, err := fetchDbItemsWithSearch(searchQuery)
		if err != nil {
			logError(fmt.Errorf("error fetching items with search query '%s': %w", searchQuery, err))
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

	items, err := fetchDbItems(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := extractPathParam(r.URL.Path, "/api/items/")
	if err != nil {
		logError(fmt.Errorf("error fetching item with ID '%s': %w", itemId, err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := fetchDbItem(itemId)
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

func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := extractPathParam(r.URL.Path, "/api/items/")
	if err != nil {
		logError(fmt.Errorf("error updating item with ID '%s': %w", itemId, err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedItem Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedItem.ID = itemId
	err = updateDbItem(&updatedItem)
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

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		logError(fmt.Errorf("error creating new item: %w", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateItem(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdItem, err := createDbItem(newItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating item: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdItem)
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		logError(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	createdUser, err := createUser(newUser)
	if err != nil {
		logError(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	tokenString, err := createToken(createdUser.ID)
	if err != nil {
		logError(fmt.Errorf("signup error: error creating token for user ID %s: %w", createdUser.ID, err))
		http.Error(w, "Error creating user token", http.StatusInternalServerError)
		return
	}

	signupResponse := struct {
		Token string       `json:"token"`
		User  UserResponse `json:"user"`
	}{
		Token: tokenString,
		User: UserResponse{
			ID:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signupResponse)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		logError(fmt.Errorf("login error: invalid request body: %w", err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	logInfo(fmt.Sprintf("Attempting login for email: %s", loginUser.Email))

	user, err := fetchUserByEmail(loginUser.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logError(fmt.Errorf("login error: no user found with email %s: %w", loginUser.Email, err))
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			logError(fmt.Errorf("login error: error fetching user by email: %w", err))
			http.Error(w, "Error logging in", http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		logError(fmt.Errorf("login error: password mismatch for email %s: %w", loginUser.Email, err))
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := createToken(user.ID)
	if err != nil {
		logError(fmt.Errorf("login error: error creating token for user ID %s: %w", user.ID, err))
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	type loginResponse struct {
		Token string       `json:"token"`
		User  UserResponse `json:"user"`
	}

	w.Header().Set("Content-Type", "application/json")

	response := loginResponse{
		Token: tokenString,
		User: UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logError(fmt.Errorf("login error: error encoding login response: %w", err))
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	logInfo(fmt.Sprintf("User logged in: %s", user.Email))
}
