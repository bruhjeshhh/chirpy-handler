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
	log.Printf("we ballin")
	log.Fatal(srv.ListenAndServe())
}
