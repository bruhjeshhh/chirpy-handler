package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bruhjeshhh/chirpy/internal/database"
	"github.com/google/uuid"
)

type request struct {
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

func (cfg *apiConfig) Chirp(w http.ResponseWriter, r *http.Request) {

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

	req.Body = cleaned
	id, _ := uuid.Parse(req.UserID)

	feedback, eror := cfg.db.PostChirp(r.Context(), database.PostChirpParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      req.Body,
		UserID:    id,
	})

	if eror != nil {
		log.Println(eror)
		respondWithError(w, 400, "db ke wqt dikkat")
		return
	}

	type response struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
		UserID    string    `json:"user_id"`
	}

	resp := response{
		ID:        feedback.ID.String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      req.Body,
		UserID:    req.UserID,
	}
	respondWithJson(w, 201, resp)
}
