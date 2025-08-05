package main

import (
	"encoding/json"
	"net/http"

	"github.com/Hedonysym/go_server/internal/auth"
)

func (cfg *apiConfig) userLoginEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	login := userLogin{}
	err := decoder.Decode(&login)
	if err != nil {
		respondWithError(w, 400, "Bad Request")
		return
	}
	user, err := cfg.db.GetUserByEmail(r.Context(), login.Email)
	if err != nil {
		respondWithError(w, 401, "email not in use")
		return
	}

	err = auth.CheckPasswordHash(login.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, 401, "incorrect email or password")
		return
	}

	respondWithJSON(w, 200, userReformatter(user))
}
