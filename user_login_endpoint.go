package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/auth"
	"github.com/Hedonysym/go_server/internal/database"
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
	expTime := getExpirationInSecs(login.ExpiresInSeconds)

	token, err := auth.MakeJWT(user.ID, cfg.secret, expTime)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("token generation error: %v", err))
		return
	}
	refTString, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, 400, "something is wrong ig, this error should never happen")
	}

	refreshParams := database.NewRefreshTokenParams{
		Token:     refTString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		ExpiresAt: time.Now().AddDate(0, 0, 60),
	}

	_, err = cfg.db.NewRefreshToken(r.Context(), refreshParams)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("bad refresh token: %v", err))
		return
	}

	respondWithJSON(w, 200, userReformatter(user, &token, &refTString))
}
