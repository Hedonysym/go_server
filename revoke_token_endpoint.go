package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/auth"
	"github.com/Hedonysym/go_server/internal/database"
)

func (cfg *apiConfig) revokeTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 400, "token error")
		return
	}

	now := time.Now()

	revoked := sql.NullTime{
		Time:  now,
		Valid: true,
	}

	params := database.RevokeRefreshTokenParams{
		Token:     token,
		RevokedAt: revoked,
		UpdatedAt: now,
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), params)
	if err != nil {
		respondWithError(w, 400, "revoke token error")
		return
	}
	w.WriteHeader(204)
}
