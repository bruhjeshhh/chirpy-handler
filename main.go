package main

import (
	"log"
	"net/http"
	"sync/atomic"

	_ "github.com/lib/pq"
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
