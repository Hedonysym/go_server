package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

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
