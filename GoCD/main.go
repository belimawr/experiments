package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/hello/", &helloHandler{
		version: "0.1",
	})

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("Could not start the server")
	}
}
