package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/health/", health)
	http.Handle("/hello/", &helloHandler{
		version: "0.2",
	})

	if err := http.ListenAndServe("0.0.0.0:5000", nil); err != nil {
		log.Fatal("Could not start the server")
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
