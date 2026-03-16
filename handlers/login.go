package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"logIn/model"
	"logIn/services"
)

// LoginHandler godoc
// @Summary      Login a user
// @Description  Validates credentials and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User credentials"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /login [post]
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	var storedPassword string
	var userID int

	err = db.QueryRow(
		"SELECT id, password FROM users WHERE email = ?",
		user.Email,
	).Scan(&userID, &storedPassword)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !services.CheckPassword(storedPassword, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := services.GenerateToken(userID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
}