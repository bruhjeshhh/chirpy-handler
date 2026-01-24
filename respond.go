package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

func respondWithJson(w http.ResponseWriter, n int, payload any) {

	resp, eror := json.Marshal(payload)
	if eror != nil {
		log.Printf("idhar dikkat aai3")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(n)
	w.Write(resp)

}

func cleanseBody(s string) string {
	splits := strings.Split(s, " ")
	for i, str := range splits {
		lowstr := strings.ToLower(str)
		if lowstr == "kerfuffle" || lowstr == "sharbert" || lowstr == "fornax" {
			splits[i] = "****"
		}
	}
	return strings.Join(splits, " ")

}
