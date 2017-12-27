package handlers

import (
	"net/http"

	"github.com/dimiro1/health"
)

// NewHealthHandler returns a health handler
func NewHealthHandler(version string) http.Handler {
	handler := health.NewHandler()
	handler.AddInfo("version", version)

	return handler
}
