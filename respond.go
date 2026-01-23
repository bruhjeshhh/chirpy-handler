package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorstc struct {
		Error string `json:"error"`
	}
	errormsg := errorstc{
		Error: msg,
	}

	resp, eror := json.Marshal(errormsg)
	if eror != nil {
		log.Printf("idhar dikkat aai")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)

}

func respondWithJson(w http.ResponseWriter) {
	type validity struct {
		Valid bool `json:"valid"`
	}

	vald := validity{
		Valid: true,
	}

	resp, eror := json.Marshal(vald)
	if eror != nil {
		log.Printf("idhar dikkat aai3")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resp)

}
