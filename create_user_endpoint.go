package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/auth"
	"github.com/Hedonysym/go_server/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	login := userLogin{}
	err := decoder.Decode(&login)
	if err != nil {
		respondWithError(w, 400, "Bad Request")
		return
	}

	hash, err := auth.HashPassword(login.Password)
	if err != nil {
		respondWithError(w, 403, "invalid password")
		return
	}

	params := database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Email:          login.Email,
		HashedPassword: hash,
	}

	user, err := cfg.db.CreateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Databse error: %v", err))
		return
	}

	respondWithJSON(w, 201, User{
		Id:         user.ID,
		Created_at: user.CreatedAt,
		Updated_at: user.UpdatedAt,
		Email:      user.Email,
	})
}
