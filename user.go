package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/bruhjeshhh/chirpy/internal/database"
)

func (cfg *apiConfig) addUser(w http.ResponseWriter, r *http.Request) {
	type emailrecv struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	eml := emailrecv{}
	err := decoder.Decode(&eml)
	if err != nil {
		respondWithError(w, 400, "something went wrong")
		return
	}
	email := eml.Email

	_, errr := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     email,
	},
	)

	if errr != nil {
		respondWithError(w, 400, "db ke wqt dikkat")
		return
	}

	type respnse struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	resp := respnse{
		ID:        "Sfdsf",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email,
	}
	respondWithJson(w, 201, resp)

}
