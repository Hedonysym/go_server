package main

import (
	"sync/atomic"
	"time"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Id            uuid.UUID `json:"id"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	Email         string    `json:"email"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	Is_chirpy_red bool      `json:"is_chirpy_red"`
}

type Chirp struct {
	Id         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Body       string    `json:"body"`
	User_id    uuid.UUID `json:"user_id"`
}

type apiConfig struct {
	fileserverhits atomic.Int32
	db             *database.Queries
	platform       string
	secret         string
	polka_key      string
}

type userLogin struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	ExpiresInSeconds int    `json:"expires_in_seconds"`
}
