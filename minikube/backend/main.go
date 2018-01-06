package main

import (
	"fmt"
	stdlog "log"
	"net/http"

	"github.com/belimawr/experiments/minikube/backend/config"
	"github.com/belimawr/experiments/minikube/backend/handlers"
	"github.com/caarlos0/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		stdlog.Fatalf("error parsing config: %q", err)
	}

	logger := cfg.ZeroLog()
	stdlog.SetFlags(0)
	stdlog.SetOutput(logger)

	health := handlers.NewHealthHandler("Development")
	hello := handlers.NewHelloHandler("who")

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Use(hlog.NewHandler(logger))
	router.Use(handlers.NewLogMiddleware())

	router.Method("GET", "/health", health)
	router.Method("GET", "/hello", hello)

	addr := fmt.Sprintf(":%d", cfg.Port)

	logger.Info().Msgf("starting webserver on: %q", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Panic().Err(err)
	}
}
