package middleware

import (
	"net/http"

	"github.com/klimenkokayot/avito-go/libs/logger"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.WriteHeader(code)
}

func LoggerMiddleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		logger.Debug("Request started")
		logger.Debug("Request done")
		return h
	}
}
