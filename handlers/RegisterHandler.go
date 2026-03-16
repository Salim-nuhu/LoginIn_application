package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"logIn/model"
	"logIn/services"
)
// RegisterHandler godoc
// @Summary      Register a new user
// @Description  Creates a new user with a hashed password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User credentials"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /register [post]
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user model.User

		// 1. Decode FIRST — populate user from request body
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		// 2. Validate AFTER — now user.Email and user.Password actually have values
		if user.Email == "" || user.Password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		// 3. Hash the password
		hashedPassword, err := services.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// 4. Save to database
		query := "INSERT INTO users(email, password) VALUES (?, ?)"
		_, err = db.Exec(query, user.Email, hashedPassword)
		if err != nil {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
	}
}