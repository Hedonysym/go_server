package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) allChirpsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	authId := r.URL.Query().Get("author_id")
	ifsort := r.URL.Query().Get("sort")
	if authId != "" {
		id, err := uuid.Parse(authId)
		if err != nil {
			respondWithError(w, 401, "invalid id")
		}
		chirps, err := cfg.db.GetAllChirpsByUser(ctx, id)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("database error:  %v", err))
			return
		}
		newchirps := []Chirp{}
		for _, chirp := range chirps {
			newchirps = append(newchirps, chirpReformatter(chirp))
		}
		if ifsort == "desc" {
			sort.Slice(newchirps, func(i, j int) bool { return newchirps[j].Created_at.Before(newchirps[i].Created_at) })
		}
		respondWithJSON(w, 200, newchirps)
		return

	} else {
		chirps, err := cfg.db.AllChirps(ctx)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("database error:  %v", err))
			return
		}
		newchirps := []Chirp{}
		for _, chirp := range chirps {
			newchirps = append(newchirps, chirpReformatter(chirp))
		}
		if ifsort == "desc" {
			sort.Slice(newchirps, func(i, j int) bool { return newchirps[j].Created_at.Before(newchirps[i].Created_at) })
		}
		respondWithJSON(w, 200, newchirps)
	}
}
