package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/belimawr/experiments/minikube/backend/config"
	"github.com/belimawr/experiments/minikube/backend/handlers"
	"github.com/caarlos0/env"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error parsing config: %q", err)
	}

	health := handlers.NewHealthHandler("Development")

	http.Handle("/health/", handlers.NewLogMiddleware(health))

	addr := fmt.Sprintf(":%d", cfg.Port)

	log.Printf("starting webserver on: %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Panic(err)
	}
}
