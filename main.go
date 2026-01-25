package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/bruhjeshhh/chirpy/internal/database"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

//	type Server struct {
//		Addr string
//		Handler Handler
//	}

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbURL := os.Getenv("DB_URL")

	dbz, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbz.Close()
	dbQueries := database.New(dbz)

	ptr := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: ptr,
	}
	var cfg apiConfig
	cfg.db = dbQueries
	wrappedHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	ptr.Handle("/app/", cfg.middlewareMetricsInc(wrappedHandler))
	// ptr.Handle("/metrics", cfg.middlewareMetricsInc(http.HandlerFunc(cfg.fetchmetric)))

	ptr.HandleFunc("GET /api/healthz", app)
	ptr.HandleFunc("POST /admin/reset", cfg.resetmetric)
	ptr.HandleFunc("GET /admin/metrics", cfg.fetchmetric)
	ptr.HandleFunc("POST /api/chirps", cfg.Chirp)
	ptr.HandleFunc("GET /api/chirps", cfg.fetchChirps)
	ptr.HandleFunc("GET /api/chirps/{chirpID}", cfg.fetchChirpsbyID)

	ptr.HandleFunc("POST /api/users", cfg.addUser)
	ptr.HandleFunc("POST /api/login", cfg.loginUser)

	log.Printf("we ballin")
	log.Fatal(srv.ListenAndServe())
}
