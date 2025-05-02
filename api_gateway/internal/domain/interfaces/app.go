package interfaces

import (
	"context"
)

type Application interface {
	// Жизненный цикл
	Run() error
	Shutdown(ctx context.Context) error

	// Мониторинг
	Metrics() (metrics map[string]interface{}, err error)
}
