package main

import (
	"log"
	"net/http"
	"testing"

	"net/http/httptest"
)

func Test_hello(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/hello/", nil)

	h := &helloHandler{
		version: "test",
	}

	h.ServeHTTP(w, r)

	if w.Code != 200 {
		log.Fatalf("w.Code = %d, wanted 200", w.Code)
	}
}

func Test_hello_not_GET(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/hello/", nil)

	h := &helloHandler{
		version: "test",
	}

	h.ServeHTTP(w, r)

	if w.Code != http.StatusMethodNotAllowed {
		log.Fatalf("w.Code = %d, wanted %d", w.Code, http.StatusMethodNotAllowed)
	}
}
