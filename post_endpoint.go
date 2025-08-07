package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hedonysym/go_server/internal/auth"
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
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("something happened: %v", err))
		return
	}
	id, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("invalid auth token: %v", err))
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
		UserID:    id,
	}
	chirp, err := cfg.db.PostChirp(r.Context(), params)
	if err != nil {
		respondWithError(w, 400, "error posting chirp")
		return
	}
	respondWithJSON(w, 201, Chirp{
		Id:         chirp.ID,
		Created_at: chirp.CreatedAt,
		Updated_at: chirp.UpdatedAt,
		Body:       chirp.Body,
		User_id:    chirp.UserID,
	})

}
