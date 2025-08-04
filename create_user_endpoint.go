package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type userEmail struct {
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	email := userEmail{}
	err := decoder.Decode(&email)
	if err != nil {
		respondWithError(w, 400, "Bad Request")
		return
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email.Email,
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
