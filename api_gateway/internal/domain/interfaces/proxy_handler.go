package interfaces

import "net/http"

type ProxyHandler interface {
	// Точка маршрутизации
	ServeHTTP(w http.ResponseWriter, r *http.Request)

	// Микросервисы
	ProxyAuthService(w http.ResponseWriter, r *http.Request)

	// Middleware
	AuthMiddleware(next http.Handler) http.Handler
	LoggerMiddleware(next http.Handler) http.Handler

	// Мониторинг
	Metrics() (metrics map[string]interface{}, err error)
}
