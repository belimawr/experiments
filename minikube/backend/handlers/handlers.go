package handlers

import (
	"fmt"
	"net/http"

	"github.com/belimawr/uhttp"
)

// NewHelloHandler returns a http.Handler that says hello
func NewHelloHandler(key string) http.Handler {
	return uhttp.WrapFunc(func(w http.ResponseWriter, r *http.Request) error {
		who := r.URL.Query().Get(key)

		if who == "" {
			who = "World"
		}

		fmt.Fprintf(w, "Hello %s!", who)

		return nil
	})
}
