package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/database"
)

func chirpReformatter(dat database.Chirp) Chirp {
	return Chirp{
		Id:         dat.ID,
		Created_at: dat.CreatedAt,
		Updated_at: dat.UpdatedAt,
		Body:       dat.Body,
		User_id:    dat.UserID,
	}
}

func (cfg *apiConfig) allChirpsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	chirps, err := cfg.db.AllChirps(ctx)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("database error:  %v", err))
		return
	}
	newchirps := []Chirp{}
	for _, chirp := range chirps {
		newchirps = append(newchirps, chirpReformatter(chirp))
	}
	respondWithJSON(w, 200, newchirps)
}
