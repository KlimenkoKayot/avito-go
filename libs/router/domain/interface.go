package domain

import (
	"net/http"
)

type Router interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	Use(middleware MiddlewareFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type (
	Handler        = http.Handler
	HandlerFunc    = func(w http.ResponseWriter, r *http.Request)
	MiddlewareFunc = func(next Handler) Handler
)
