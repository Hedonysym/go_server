package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/google/uuid"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type responseParams struct {
		Error string `json:"error"`
	}
	res := responseParams{
		Error: msg,
	}
	w.Header().Set("Content_Type", "application/json")
	jsonData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"JSON marshalling error"}`))
		return
	}
	w.WriteHeader(code)
	w.Write(jsonData)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content_type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func profanityScrubber(msg string) string {
	words := strings.Split(msg, " ")
	for i, word := range words {
		chk := strings.ToLower(word)
		if chk == "kerfuffle" || chk == "sharbert" || chk == "fornax" {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}

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
