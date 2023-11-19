package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	jsonData, marshalError := json.Marshal(payload)
	if marshalError != nil {
		log.Printf("Error marshalling jsonData: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/jsonData")
	w.WriteHeader(code)
	write, writeError := w.Write(jsonData)
	if writeError != nil {
		log.Printf("Error writing jsonData: %v\n", write)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("Responding with 5XX error: %v\n", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, errResponse{Error: msg})
}
