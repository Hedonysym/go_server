package main

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	hits := apiCfg.fileserverhits.Load()
	response := fmt.Sprintf(
		`<html>
  			<body>
   				<h1>Welcome, Chirpy Admin</h1>
    			<p>Chirpy has been visited %d times!</p>
  			</body>
		</html>`,
		hits)
	_, err := w.Write([]byte(response))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (apiCfg *apiConfig) metricsResetHandler(w http.ResponseWriter, r *http.Request) {
	apiCfg.fileserverhits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hits reset to 0"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
