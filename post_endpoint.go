package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) postEndpointHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		Userid uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	req := parameters{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}
	if len(req.Body) > 140 {
		respondWithError(w, 400, "Chirp too long")
		return
	}
	cleaned := profanityScrubber(req.Body)
	params := database.PostChirpParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Body:      cleaned,
		UserID:    req.Userid,
	}
	chirp, err := cfg.db.PostChirp(r.Context(), params)
	respondWithJSON(w, 201, Chirp{
		Id:         chirp.ID,
		Created_at: chirp.CreatedAt,
		Updated_at: chirp.UpdatedAt,
		Body:       chirp.Body,
		User_id:    chirp.UserID,
	})
}
