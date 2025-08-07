package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("postgress load failed")
		os.Exit(1)
	}
	dbQueries := database.New(db)
	mux := http.NewServeMux()
	cfg := &apiConfig{
		db:       dbQueries,
		platform: platform,
		secret:   secret,
	}

	mux.HandleFunc("GET /api/healthz", readyEndpointHandler)
	mux.HandleFunc("GET /admin/metrics", cfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", cfg.metricsResetHandler)
	mux.HandleFunc("POST /api/chirps", cfg.postEndpointHandler)
	mux.HandleFunc("POST /api/users", cfg.createUserEndpoint)
	mux.HandleFunc("GET /api/chirps", cfg.allChirpsEndpoint)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.getChirpEndpoint)
	mux.HandleFunc("POST /api/login", cfg.userLoginEndpoint)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/app/", cfg.middlewareMetricsInc(
		http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
