package main

import (
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

	ptr.HandleFunc("/healthz", app)
	ptr.HandleFunc("/reset", cfg.resetmetric)
	ptr.HandleFunc("/metrics", cfg.fetchmetric)

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
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	hits := cfg.fileserverHits.Load()
	resp := fmt.Sprintf("Hits: %d", hits)
	w.Write([]byte(resp))

}
