package router

import (
	"fmt"

	"github.com/klimenkokayot/avito-go/libs/router/adapters/gorilla"
	"github.com/klimenkokayot/avito-go/libs/router/domain"
)

const (
	AdapterGorilla = "gorilla"
)

var (
	ErrUnknownAdapter = fmt.Errorf("роутер не поддерижвается")
)

type (
	Router = domain.Router
)

type Config struct {
	Name string
}

func NewAdapter(cfg *Config) (Router, error) {
	switch cfg.Name {
	case AdapterGorilla:
		return gorilla.NewAdapter()
	default:
		return nil, ErrUnknownAdapter
	}
}
