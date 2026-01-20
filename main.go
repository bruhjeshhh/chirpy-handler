package main

import (
	"log"
	"net/http"
)

//	type Server struct {
//		Addr string
//		Handler Handler
//	}
func main() {
	ptr := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: ptr,
	}
	wrappedHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	ptr.Handle("/app/", wrappedHandler)
	ptr.HandleFunc("/healthz", app)

	log.Printf("we ballin")
	log.Fatal(srv.ListenAndServe())
}

func app(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
