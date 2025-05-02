package interfaces

import (
	"context"
	"net/http"
)

type ProxyServer interface {
	// Жизненный цикл
	Run() error
	Shutdown(ctx context.Context) error

	// Маршрутизация
	RegisterRoute(method, path string, route http.HandlerFunc) error
	RegisterMiddleware(middleware func(http.Handler) http.Handler) error

	// Middleware (серверные)
	RateLimitMiddleware(next http.Handler) http.Handler

	// Мониторинг
	Metrics() (metrics map[string]interface{}, err error)
}
