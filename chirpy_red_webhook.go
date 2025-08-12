package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) chirpyRedWebhook(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Event string `json:"event"`
		Data  struct {
			User_id uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	req := Request{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if req.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	err = cfg.db.UpgradeUserToRed(r.Context(), req.Data.User_id)
	if err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(204)
}
