package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Hedonysym/go_server/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverhits atomic.Int32
	db             *database.Queries
	platform       string
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverhits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
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
	}

	mux.HandleFunc("GET /api/healthz", readyEndpointHandler)
	mux.HandleFunc("GET /admin/metrics", cfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", cfg.metricsResetHandler)
	mux.HandleFunc("POST /api/validate_chirp", postEndpointHandler)
	mux.HandleFunc("POST /api/users", cfg.createUserEndpoint)

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
