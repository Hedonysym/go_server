package main

import (
	"encoding/json"
	"net/http"
)

func postEndpointHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	r.Header.Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Something went wrong"}`))
		return
	}
	if len(params.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Chirp is too long"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"valid":true}`))
}
