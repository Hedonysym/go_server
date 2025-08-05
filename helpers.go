package main

import (
	"encoding/json"
	"net/http"
	"strings"

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

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverhits.Add(1)
		next.ServeHTTP(w, r)
	})
}

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
