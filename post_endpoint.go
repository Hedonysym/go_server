package main

import (
	"encoding/json"
	"net/http"
	"strings"
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

func postEndpointHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp too long")
		return
	}
	type Response struct {
		Cleaned string `json:"cleaned_body"`
	}
	cleaned := profanityScrubber(params.Body)
	respondWithJSON(w, 200, Response{Cleaned: cleaned})
}
