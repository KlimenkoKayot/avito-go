package logrus

import (
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/sirupsen/logrus"
)

type LogrusAdapter struct {
	*logrus.Logger
}

func (a *LogrusAdapter) Debug(msg string, fields ...logger.Field) {
	a.Logger.WithFields(toLogrusFields(fields)).Debug(msg)
}

func (a *LogrusAdapter) Error(msg string, fields ...logger.Field) {
	a.Logger.WithFields(toLogrusFields(fields)).Error(msg)
}

func (a *LogrusAdapter) Fatal(msg string, fields ...logger.Field) {
	a.Logger.WithFields(toLogrusFields(fields)).Fatal(msg)
}

func (a *LogrusAdapter) Info(msg string, fields ...logger.Field) {
	a.Logger.WithFields(toLogrusFields(fields)).Info(msg)
}

func (a *LogrusAdapter) Warn(msg string, fields ...logger.Field) {
	a.Logger.WithFields(toLogrusFields(fields)).Warn(msg)
}

func NewAdapter(level logger.Level) (logger.Logger, error) {
	logrusLogger := logrus.New()
	adapter := &LogrusAdapter{
		logrusLogger,
	}
	logrusLogger.SetLevel(logrus.Level(level))
	return adapter, nil
}

func toLogrusFields(fields []logger.Field) logrus.Fields {
	converted := logrus.Fields{}
	for _, val := range fields {
		converted[val.Key] = val.Value
	}
	return converted
}
