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
		utils.LogError(fmt.Errorf("sign up decode error: %w", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError(fmt.Errorf("generate hash error: %w", err))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	userRepo := data.NewUserRepository(deps.DB)
	inventoryRepo := data.NewInventoryRepository(deps.DB)

	inventories, err := inventoryRepo.List()
	if err != nil {
		utils.LogError(fmt.Errorf("get inventories error: %w", err))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// TODO: Refactor if signup/create user flow is enhanced and specifies which inventories a user has access to
	createdUser, err := userRepo.CreateUser(newUser, inventories)
	if err != nil {
		utils.LogError(fmt.Errorf("create user error: %w", err))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	tokenString, err := auth.CreateToken(createdUser.ID, createdUser.Role, deps.JwtSecret)
	if err != nil {
		utils.LogError(fmt.Errorf("create token for user error: user ID %s: %w", createdUser.ID, err))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	response := SignupResponse{
		Token: tokenString,
		User: UserResponse{
			ID:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
	}

	err = utils.WriteJSONResponse(w, http.StatusCreated, response, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("sign up json response error: %w", err))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
}

func (deps *HandlerDependencies) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		utils.LogError(fmt.Errorf("login decode error: %w", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	utils.LogInfo(fmt.Sprintf("Attempting login for email: %s", loginUser.Email))

	repo := data.NewUserRepository(deps.DB)

	user, err := repo.FetchUserByEmail(loginUser.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.LogError(fmt.Errorf("login no user found error: email %s: %w", loginUser.Email, err))
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			utils.LogError(fmt.Errorf("login error: error fetching user by email: %w", err))
			http.Error(w, "Error logging in", http.StatusInternalServerError)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		utils.LogError(fmt.Errorf("login error: password mismatch for email %s: %w", loginUser.Email, err))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.CreateToken(user.ID, user.Role, deps.JwtSecret)

	if err != nil {
		utils.LogError(fmt.Errorf("login error: error creating token for user ID %s: %w", user.ID, err))
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Token: tokenString,
		User: UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}

	err = utils.WriteJSONResponse(w, http.StatusOK, response, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("login error: error encoding login response: %w", err))
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	utils.LogInfo(fmt.Sprintf("User logged in: %s", user.Email))
}

func (deps *HandlerDependencies) HandleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	repo := data.NewUserRepository(deps.DB)

	var passwordResetReq PasswordResetRequest
	err := json.NewDecoder(r.Body).Decode(&passwordResetReq)
	if err != nil {
		utils.LogError(fmt.Errorf("update password decode error: %w", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userID, _, err := auth.AuthenticateUser(r, deps.JwtSecret)
	if err != nil {
		utils.LogError(fmt.Errorf("update password token auth error: %w", err))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := repo.FetchUserByID(userID)
	if err != nil {
		utils.LogError(fmt.Errorf("update password user not found error: %w", err))
		http.Error(w, "Update password error", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordResetReq.CurrentPassword))
	if err != nil {
		utils.LogError(fmt.Errorf("update password comparison error: %w", err))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordResetReq.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError(fmt.Errorf("update password new password hash error: %w", err))
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	err = repo.UpdatePassword(user.ID, string(hashedPassword))
	if err != nil {
		utils.LogError(fmt.Errorf("update password database error: %w", err))
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Password updated successfully"}

	err = utils.WriteJSONResponse(w, http.StatusOK, response, nil)
	if err != nil {
		utils.LogError(fmt.Errorf("account update error: error encoding update password response: %w", err))
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}
}
