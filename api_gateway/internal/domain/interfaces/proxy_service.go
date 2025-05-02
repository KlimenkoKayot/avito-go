package interfaces

import (
	"context"
	"net/http"

	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model"
)

type ProxyService interface {
	// Проксирование
	ForwardRequest(ctx context.Context, r *http.Request) (resp *http.Response, err error)

	// Регистрация микросервисов
	RegisterService(name string, config *model.ServiceConfig) error
	GetAvailableService() []string

	// Обработка ошибок проксирования
	HandleProxyError(w http.ResponseWriter, err error) error

	// Мониторинг
	Metrics() (metrics map[string]interface{}, err error)
}
