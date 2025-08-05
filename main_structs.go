package main

import (
	"sync/atomic"
	"time"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Email      string    `json:"email"`
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
}

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
