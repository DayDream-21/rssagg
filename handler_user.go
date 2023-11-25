package main

import (
	"encoding/json"
	"fmt"
	"github.com/DayDream-21/rssagg/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	prams := parameters{}
	decodeError := decoder.Decode(&prams)
	if decodeError != nil {
		responseWithJson(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s\n", decodeError))
		return
	}

	user, userCreateError := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      prams.Name,
	})
	if userCreateError != nil {
		responseWithJson(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %s\n", userCreateError))
		return
	}

	responseWithJson(w, http.StatusOK, user)
}
