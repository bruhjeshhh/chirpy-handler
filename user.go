package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/bruhjeshhh/chirpy/internal/auth"
	"github.com/bruhjeshhh/chirpy/internal/database"
)

type emailrecv struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type respnse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) addUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	eml := emailrecv{}
	err := decoder.Decode(&eml)
	if err != nil {
		respondWithError(w, 400, "something went wrong")
		return
	}
	email := eml.Email
	hashit, hasher := auth.HashPassword(eml.Password)
	if hasher != nil {
		respondWithError(w, 400, "hashing went")
		return
	}

	feedback, errr := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Email:      email,
		HashedPswd: hashit,
	},
	)

	if errr != nil {
		respondWithError(w, 400, "db ke wqt dikkat")
		return
	}
	id := feedback.ID

	resp := respnse{
		ID:        id.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email,
	}
	respondWithJson(w, 201, resp)

}

func (cfg *apiConfig) loginUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	eml := emailrecv{}
	err := decoder.Decode(&eml)
	if err != nil {
		respondWithError(w, 400, "something went wrong")
		return
	}
	usermail := eml.Email
	hashedpswdL, herr := cfg.db.GetHashedPswd(r.Context(), usermail)
	if herr != nil {
		respondWithError(w, 400, "not found")
		return
	}
	match, matcherr := auth.CheckPasswordHash(eml.Password, hashedpswdL.HashedPswd)
	if matcherr != nil {
		log.Println(matcherr)
		respondWithError(w, 400, "something went wrong here")
		return
	}
	if match == false {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	resp := respnse{
		ID:        hashedpswdL.ID.String(),
		CreatedAt: hashedpswdL.CreatedAt,
		UpdatedAt: hashedpswdL.UpdatedAt,
		Email:     hashedpswdL.Email,
	}
	respondWithJson(w, 200, resp)

}
