package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

//	type Server struct {
//		Addr string
//		Handler Handler
//	}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	ptr := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: ptr,
	}
	var cfg apiConfig
	wrappedHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	ptr.Handle("/app/", cfg.middlewareMetricsInc(wrappedHandler))
	// ptr.Handle("/metrics", cfg.middlewareMetricsInc(http.HandlerFunc(cfg.fetchmetric)))

	ptr.HandleFunc("GET /api/healthz", app)
	ptr.HandleFunc("POST /admin/reset", cfg.resetmetric)
	ptr.HandleFunc("GET /admin/metrics", cfg.fetchmetric)
	ptr.HandleFunc("POST /api/validate_chirp", validChirp)

	log.Printf("we ballin")
	log.Fatal(srv.ListenAndServe())
}

func app(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) resetmetric(w http.ResponseWriter, r *http.Request) {
	hits := cfg.fileserverHits.Load()
	cfg.fileserverHits.Add(-hits)
}

func (cfg *apiConfig) fetchmetric(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	hits := cfg.fileserverHits.Load()
	resp := fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1> <p>Chirpy has been visited %d times!</p></body></html>", hits)
	w.Write([]byte(resp))

}

func validChirp(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	req := request{}

	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "somethingwentwrong")
		return
	}

	if len(req.Body) > 140 {
		respondWithError(w, 400, "too long buddy(thats what she said)")
		return
	}
	respondWithJson(w)
}
