package main

import (
	"encoding/json"
	"net/http"

	"github.com/Hedonysym/go_server/internal/auth"
	"github.com/Hedonysym/go_server/internal/database"
)

func (cfg *apiConfig) updateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tknStr, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "token header error")
		return
	}

	userId, err := auth.ValidateJWT(tknStr, cfg.secret)
	if err != nil {
		respondWithError(w, 401, "invalid token error")
		return
	}

	reqBody := userLogin{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&reqBody)
	if err != nil {
		respondWithError(w, 401, "invalid request")
		return
	}

	pword, err := auth.HashPassword(reqBody.Password)
	if err != nil {
		respondWithError(w, 401, "password hashed incorrectly, check password")
		return
	}

	params := database.UpdateUserParams{
		ID:             userId,
		Email:          reqBody.Email,
		HashedPassword: pword,
	}
	user, err := cfg.db.UpdateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, 401, "error updating user info")
		return
	}

	respondWithJSON(w, 200, userReformatter(user, nil, nil))
}
