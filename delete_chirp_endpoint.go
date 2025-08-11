package main

import (
	"net/http"

	"github.com/Hedonysym/go_server/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) deleteChirpEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	tknStr, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "invalid token header")
		return
	}
	userId, err := auth.ValidateJWT(tknStr, cfg.secret)
	if err != nil {
		respondWithError(w, 401, "invalid token string")
		return
	}

	chirpId, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, 400, "Something heppened")
	}
	chirp, err := cfg.db.GetChirpByChirpId(r.Context(), chirpId)
	if err != nil {
		respondWithError(w, 404, "Chirp not found")
	}
	if userId != chirp.UserID {
		respondWithError(w, 403, "thats not your chirp")
		return
	}

	err = cfg.db.DeleteChirp(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, 400, "deletion failed")
		return
	}

	w.WriteHeader(204)
}
