package zap

import (
	"fmt"

	"github.com/klimenkokayot/avito-go/libs/logger"
	"go.uber.org/zap"
)

type ZapAdapter struct {
	*zap.Logger
}

// Debug implements logger.Logger.
// Subtle: this method shadows the method (*Logger).Debug of ZapAdapter.Logger.
func (z *ZapAdapter) Debug(msg string, fields ...logger.Field) {
	z.Logger.Debug(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Error(msg string, fields ...logger.Field) {
	z.Logger.Error(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Fatal(msg string, fields ...logger.Field) {
	z.Logger.Fatal(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Info(msg string, fields ...logger.Field) {
	z.Logger.Info(msg, toZapFields(fields)...)
}

func (z *ZapAdapter) Warn(msg string, fields ...logger.Field) {
	z.Logger.Warn(msg, toZapFields(fields)...)
}

func NewAdapter(isDebug bool) (logger.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	if isDebug {
		zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	zapLogger, err := zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("%w: %s.", ErrZapBuild, err.Error())
	}
	return &ZapAdapter{
		zapLogger,
	}, nil
}

func toZapFields(fields []logger.Field) []zap.Field {
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
