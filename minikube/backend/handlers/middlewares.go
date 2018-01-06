package handlers

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/hlog"
)

// NewLogMiddleware wraps a http.Handler with a logger
func NewLogMiddleware() func(http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("remote_addr", r.RemoteAddr).
			Str("user_agent", r.UserAgent()).
			Int("status", status).
			Int("bytes", size).
			Dur("duration", duration).
			Msg("")
	})
}
