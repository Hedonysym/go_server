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
	dbQ            *database.Queries
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
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Errorf("postgress load failed")
		os.Exit(1)
	}
	dbQueries := database.New(db)
	mux := http.NewServeMux()
	apiCfg := &apiConfig{
		dbQ: dbQueries,
	}

	mux.HandleFunc("GET /api/healthz", readyEndpointHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.metricsResetHandler)
	mux.HandleFunc("POST /api/validate_chirp", postEndpointHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(
		http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
