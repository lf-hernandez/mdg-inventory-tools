package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/lf-hernandez/mdg-inventory-tools/api/auth"
	"github.com/lf-hernandez/mdg-inventory-tools/api/data"
	"github.com/lf-hernandez/mdg-inventory-tools/api/models"
	"github.com/lf-hernandez/mdg-inventory-tools/api/utils"
)

func (deps *HandlerDependencies) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	repo := data.NewUserRepository(deps.DB)

	createdUser, err := repo.CreateUser(deps.DB, newUser)
	if err != nil {
		utils.LogError(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	tokenString, err := auth.CreateToken(createdUser.ID, deps.JwtSecret)
	if err != nil {
		utils.LogError(fmt.Errorf("signup error: error creating token for user ID %s: %w", createdUser.ID, err))
		http.Error(w, "Error creating user token", http.StatusInternalServerError)
		return
	}

	signupResponse := SignupResponse{
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

func (deps *HandlerDependencies) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		utils.LogError(fmt.Errorf("login error: invalid request body: %w", err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	utils.LogInfo(fmt.Sprintf("Attempting login for email: %s", loginUser.Email))

	repo := data.NewUserRepository(deps.DB)

	user, err := repo.FetchUserByEmail(deps.DB, loginUser.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.LogError(fmt.Errorf("login error: no user found with email %s: %w", loginUser.Email, err))
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			utils.LogError(fmt.Errorf("login error: error fetching user by email: %w", err))
			http.Error(w, "Error logging in", http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		utils.LogError(fmt.Errorf("login error: password mismatch for email %s: %w", loginUser.Email, err))
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.CreateToken(user.ID, deps.JwtSecret)

	if err != nil {
		utils.LogError(fmt.Errorf("login error: error creating token for user ID %s: %w", user.ID, err))
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := LoginResponse{
		Token: tokenString,
		User: UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.LogError(fmt.Errorf("login error: error encoding login response: %w", err))
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	utils.LogInfo(fmt.Sprintf("User logged in: %s", user.Email))
}

func (deps *HandlerDependencies) HandleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repo := data.NewUserRepository(deps.DB)

	var passwordResetReq PasswordResetRequest
	err := json.NewDecoder(r.Body).Decode(&passwordResetReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := auth.AuthenticateUser(r, deps.JwtSecret)
	if err != nil {
		http.Error(w, "Unauthorized - token invalid", http.StatusUnauthorized)
		return
	}

	user, err := repo.FetchUserByID(deps.DB, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordResetReq.CurrentPassword))
	if err != nil {
		http.Error(w, "Invalid current password", http.StatusUnauthorized)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordResetReq.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError(err)
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	err = repo.UpdatePassword(deps.DB, user.ID, string(hashedPassword))
	if err != nil {
		utils.LogError(err)
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}
