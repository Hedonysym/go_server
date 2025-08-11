package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/auth"
)

func (cfg *apiConfig) refreshEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	refTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("token error: %v", err))
		return
	}

	reftoken, err := cfg.db.GetRefreshToken(r.Context(), refTokenString)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("token get error: %v", err))
		return
	}

	if reftoken.ExpiresAt.Before(time.Now()) || reftoken.RevokedAt.Valid {
		respondWithError(w, 401, "token expired/revoked, you must login again")
		return
	}

	user, err := cfg.db.GetUserByRefreshToken(r.Context(), reftoken.Token)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("user fetch error: %v", err))
		return
	}

	newtoken, err := auth.MakeJWT(user.ID, cfg.secret, getExpirationInSecs(3600))
	if err != nil {
		respondWithError(w, 400, "make token error")
		return
	}

	type Response struct {
		Token string `json:"token"`
	}

	respondWithJSON(w, 200, Response{Token: newtoken})
}
