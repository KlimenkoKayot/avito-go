package logger

import (
	"fmt"

	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/libs/logger/adapters/logrus"
	"github.com/klimenkokayot/avito-go/libs/logger/adapters/zap"
)

const (
	AdapterZap    = "zap"
	AdapterLogrus = "logrus"
)

var (
	ErrUnknownAdapter = fmt.Errorf("логгер не поддерживается")
)

type Config struct {
	Adapter string
	Level   logger.Level
}

func NewAdapter(config *Config) (logger.Logger, error) {
	switch config.Adapter {
	case AdapterZap:
		return zap.NewAdapter(config.Level)
	case AdapterLogrus:
		return logrus.NewAdapter(config.Level)
	default:
		return nil, ErrUnknownAdapter
	}
}
