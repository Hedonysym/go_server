package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirpEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	id, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, 400, "Something heppened")
	}
	chirp, err := cfg.db.GetChirpByChirpId(r.Context(), id)
	if err != nil {
		respondWithError(w, 404, "Chirp not found")
	}
	respondWithJSON(w, 200, chirpReformatter(chirp))
}
