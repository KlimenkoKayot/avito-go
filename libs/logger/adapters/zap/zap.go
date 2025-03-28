package zap

import (
	"fmt"

	"github.com/klimenkokayot/avito-go/libs/logger/domain"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapAdapter struct {
	*zap.Logger
}

// Debug implements domain.Logger.
// Subtle: this method shadows the method (*Logger).Debug of ZapAdapter.Logger.
func (z *ZapAdapter) Debug(msg string, fields ...domain.Field) {
	z.Logger.Debug(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Error(msg string, fields ...domain.Field) {
	z.Logger.Error(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Fatal(msg string, fields ...domain.Field) {
	z.Logger.Fatal(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Info(msg string, fields ...domain.Field) {
	z.Logger.Info(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Warn(msg string, fields ...domain.Field) {
	z.Logger.Warn(msg, toZapFields(fields)...)
}

func NewAdapter(level domain.Level) (domain.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(zapcore.Level(level))
	zapLogger, err := zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("%w: %s.", ErrZapBuild, err.Error())
	}
	return &ZapAdapter{
		zapLogger,
	}, nil
}

func toZapFields(fields []domain.Field) []zap.Field {
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
