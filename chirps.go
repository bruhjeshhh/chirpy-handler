package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/bruhjeshhh/chirpy/internal/auth"
	"github.com/bruhjeshhh/chirpy/internal/database"
	"github.com/google/uuid"
)

type request struct {
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) Chirp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	req := request{}

	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "something went wrong")
		return
	}

	tokunn, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "cant fetch token")
		return
	}

	id, errvald := auth.ValidateJWT(tokunn, cfg.jwtsecret)
	if errvald != nil {
		respondWithError(w, 401, "could not validate jwt")
		return
	}

	if len(req.Body) > 140 {
		respondWithError(w, 400, "too long buddy(thats what she said)")
		return
	}
	cleaned := cleanseBody(req.Body)

	req.Body = cleaned
	// id, _ := uuid.Parse(req.UserID)

	feedback, eror := cfg.db.PostChirp(r.Context(), database.PostChirpParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      req.Body,
		UserID:    id,
	})

	if eror != nil {
		// log.Println(eror)
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
		UserID:    id.String(),
	}
	respondWithJson(w, 201, resp)
}

func (cfg *apiConfig) fetchChirps(w http.ResponseWriter, r *http.Request) {
	chirps := []Chirp{}
	s := r.URL.Query().Get("author_id")
	// sorst := r.URL.Query().Get("sort")
	if s == "" {
		resp, err := cfg.db.GetChirps(r.Context())

		if err != nil {
			// log.Println(err)
			respondWithError(w, 400, "db ke wqt dikkat")
			return
		}

		for _, c := range resp {
			chirps = append(chirps, Chirp{
				CreatedAt: c.CreatedAt,
				Body:      c.Body,
			})
		}

	} else {
		uids, _ := uuid.Parse(s)
		resp, err := cfg.db.GetChirpsByAuthor(r.Context(), uids)

		if err != nil {
			// log.Println(err)
			respondWithError(w, 400, "db ke wqt dikkat")
			return
		}

		for _, c := range resp {
			chirps = append(chirps, Chirp{
				Body:      c.Body,
				CreatedAt: c.CreatedAt,
			})
		}
	}

	sortDirection := "asc"
	sortDirectionParam := r.URL.Query().Get("sort")
	if sortDirectionParam == "desc" {
		sortDirection = "desc"
	}
	sort.Slice(chirps, func(i, j int) bool {
		if sortDirection == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	respondWithJson(w, http.StatusOK, chirps)

}

func (cfg *apiConfig) fetchChirpsbyID(w http.ResponseWriter, r *http.Request) {
	chirpsID := r.PathValue("chirpID")
	id, _ := uuid.Parse(chirpsID)
	c, err := cfg.db.GetChirpsbyID(r.Context(), id)

	if err != nil {
		// log.Println(eror)
		respondWithError(w, 404, "db ke wqt dikkat")
		return
	}
	chirps := Chirp{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Body:      c.Body,
		UserID:    c.UserID,
	}

	respondWithJson(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) deleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "token not found")
		return
	}
	id, err := auth.ValidateJWT(token, cfg.jwtsecret)
	if err != nil {
		respondWithError(w, 401, "could not validate")
		return
	}
	chirpsID := r.PathValue("chirpID")
	uuidchirpid, err := uuid.Parse(chirpsID)

	chirptodeltet, err := cfg.db.GetChirpsbyID(r.Context(), uuidchirpid)
	if chirptodeltet.UserID != id {
		respondWithError(w, 403, "not authorized")
	}

	erdr := cfg.db.DeleteChirp(r.Context(), database.DeleteChirpParams{
		ID:     uuidchirpid,
		UserID: id,
	})
	if erdr != nil {
		respondWithError(w, 404, "could not delete")
		return
	}
	respondWithJson(w, 204, id)

}
