package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, http.StatusOK, struct{}{})
}
