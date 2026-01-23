package main

import (
	"encoding/json"
	"net/http"
)

func validChirp(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	req := request{}

	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "something went wrong")
		return
	}

	if len(req.Body) > 140 {
		respondWithError(w, 400, "too long buddy(thats what she said)")
		return
	}
	cleaned := cleanseBody(req.Body)
	respondWithJson(w, cleaned)
}
