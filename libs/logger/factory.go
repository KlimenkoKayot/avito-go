package logger

import (
	"fmt"

	"github.com/klimenkokayot/avito-go/libs/logger/adapters/logrus"
	"github.com/klimenkokayot/avito-go/libs/logger/adapters/zap"
	"github.com/klimenkokayot/avito-go/libs/logger/domain"
)

const (
	AdapterZap    = "zap"
	AdapterLogrus = "logrus"
)

var (
	ErrUnknownAdapter = fmt.Errorf("логгер не поддерживается")
)

type (
	Level  = domain.Level
	Logger = domain.Logger
)

type Config struct {
	Adapter string
	Level   Level
}

func NewAdapter(config *Config) (Logger, error) {
	switch config.Adapter {
	case AdapterZap:
		return zap.NewAdapter(config.Level)
	case AdapterLogrus:
		return logrus.NewAdapter(config.Level)
	default:
		return nil, ErrUnknownAdapter
	}
}
