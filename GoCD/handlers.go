package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type helloHandler struct {
	version string
}

type response struct {
	Message    string
	SomeNumber int64
	Version    string
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if r.Method != "GET" {
		http.Error(w, "Invalid method, GET required", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resp := response{
		Message:    "Hello World",
		SomeNumber: int64(rand.Int()),
		Version:    h.version,
	}

	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)

}
