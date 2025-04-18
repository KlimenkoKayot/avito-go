package zap

import (
	"fmt"

	"github.com/klimenkokayot/avito-go/libs/logger/domain"
	"github.com/klimenkokayot/avito-go/libs/logger/pkg/colorise"
	"github.com/klimenkokayot/avito-go/libs/logger/pkg/formatter"
	"go.uber.org/zap"
)

type ZapAdapter struct {
	logger    *zap.Logger
	fields    []zap.Field
	formatter *formatter.Formatter
}

func (z *ZapAdapter) WithFields(fields ...domain.Field) domain.Logger {
	zapFields := toRouterFields(fields)
	return &ZapAdapter{
		logger:    z.logger,
		fields:    zapFields,
		formatter: z.formatter,
	}
}

func (z *ZapAdapter) WithLayer(name string) domain.Logger {
	return &ZapAdapter{
		logger:    z.logger,
		fields:    z.fields,
		formatter: formatter.NewFormatter(name),
	}
}

func (z *ZapAdapter) Debug(msg string, fields ...domain.Field) {
	msg = z.formatter.FormatMessage(msg)
	zapFields := append(toRouterFields(fields), z.fields...)
	z.logger.Debug(msg, zapFields...)
}

func (z *ZapAdapter) Error(msg string, fields ...domain.Field) {
	msg = z.formatter.FormatMessage(msg)
	zapFields := append(toRouterFields(fields), z.fields...)
	colorised := colorise.ColorString(msg, colorise.ColorRed)
	z.logger.Error(colorised, zapFields...)
}

func (z *ZapAdapter) Fatal(msg string, fields ...domain.Field) {
	msg = z.formatter.FormatMessage(msg)
	zapFields := append(toRouterFields(fields), z.fields...)
	z.logger.Fatal(colorise.ColorString(msg, colorise.ColorRed), zapFields...)
}

func (z *ZapAdapter) Info(msg string, fields ...domain.Field) {
	msg = z.formatter.FormatMessage(msg)
	zapFields := append(toRouterFields(fields), z.fields...)
	z.logger.Info(msg, zapFields...)
}

func (z *ZapAdapter) Warn(msg string, fields ...domain.Field) {
	msg = z.formatter.FormatMessage(msg)
	zapFields := append(toRouterFields(fields), z.fields...)
	z.logger.Warn(colorise.ColorString(msg, colorise.ColorYellow), zapFields...)
}

func (z *ZapAdapter) OK(msg string, fields ...domain.Field) {
	msg = z.formatter.FormatMessage(msg)
	zapFields := append(toRouterFields(fields), z.fields...)
	z.logger.Info(colorise.ColorString(msg, colorise.ColorGreen), zapFields...)
}

func NewAdapter(level domain.Level) (domain.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	zapCfg.Encoding = "console"
	zapCfg.Level = toRouterLevel(level)
	zapLogger, err := zapCfg.Build()
	defer zapLogger.Sync()

	if err != nil {
		return nil, fmt.Errorf("%w: %s.", ErrZapBuild, err.Error())
	}
	return &ZapAdapter{
		logger:    zapLogger,
		fields:    make([]zap.Field, 0),
		formatter: formatter.NewFormatter(""),
	}, nil
}

func toRouterLevel(level domain.Level) zap.AtomicLevel {
	switch level {
	case domain.LevelDebug:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case domain.LevelInfo:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case domain.LevelWarn:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case domain.LevelError:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case domain.LevelFatal:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	}
}

func toRouterFields(fields []domain.Field) []zap.Field {
	converted := []zap.Field{}
	for _, val := range fields {
		field := zap.Field{}
		switch val.Value.(type) {
		case string:
			field = zap.String(val.Key, val.Value.(string))
		case int:
			field = zap.Int(val.Key, val.Value.(int))
		default:
			field = zap.Any(val.Key, val.Value)
		}
		converted = append(converted, field)
	}
	return converted
}
