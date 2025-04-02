package logrus

import (
	"github.com/klimenkokayot/avito-go/libs/logger/domain"
	"github.com/klimenkokayot/avito-go/libs/logger/pkg/colorise"
	"github.com/klimenkokayot/avito-go/libs/logger/pkg/formatter"
	"github.com/sirupsen/logrus"
)

type LogrusAdapter struct {
	logger    *logrus.Logger
	fields    logrus.Fields
	formatter *formatter.Formatter
}

func (a *LogrusAdapter) WithFields(fields ...domain.Field) domain.Logger {
	return &LogrusAdapter{
		logger:    a.logger,
		fields:    toLogrusFields(fields),
		formatter: a.formatter,
	}
}

func (a *LogrusAdapter) WithLayer(name string) domain.Logger {
	return &LogrusAdapter{
		logger:    a.logger,
		fields:    a.fields,
		formatter: formatter.NewFormatter(name),
	}
}

func (a *LogrusAdapter) Debug(msg string, fields ...domain.Field) {
	msg = a.formatter.FormatMessage(msg)
	a.logger.WithFields(a.fields).WithFields(toLogrusFields(fields)).Debug(msg)
}

func (a *LogrusAdapter) Error(msg string, fields ...domain.Field) {
	msg = a.formatter.FormatMessage(msg)
	a.logger.WithFields(a.fields).WithFields(toLogrusFields(fields)).Error(msg)
}

func (a *LogrusAdapter) Fatal(msg string, fields ...domain.Field) {
	msg = a.formatter.FormatMessage(msg)
	a.logger.WithFields(a.fields).WithFields(toLogrusFields(fields)).Fatal(colorise.ColorString(msg, colorise.ColorRed))
}

func (a *LogrusAdapter) Info(msg string, fields ...domain.Field) {
	msg = a.formatter.FormatMessage(msg)
	a.logger.WithFields(a.fields).WithFields(toLogrusFields(fields)).Info(msg)
}

func (a *LogrusAdapter) Warn(msg string, fields ...domain.Field) {
	msg = a.formatter.FormatMessage(msg)
	a.logger.WithFields(a.fields).WithFields(toLogrusFields(fields)).Warn(colorise.ColorString(msg, colorise.ColorYellow))
}

func (a *LogrusAdapter) OK(msg string, fields ...domain.Field) {
	msg = a.formatter.FormatMessage(msg)
	a.logger.WithFields(a.fields).WithFields(toLogrusFields(fields)).Warn(colorise.ColorString(msg, colorise.ColorGreen))
}

func NewAdapter(level domain.Level) (domain.Logger, error) {
	logrusLogger := logrus.New()
	adapter := &LogrusAdapter{
		logrusLogger,
		make(logrus.Fields, 0),
		formatter.NewFormatter(""),
	}
	logrusLogger.SetLevel(logrus.Level(level))
	return adapter, nil
}

func toLogrusFields(fields []domain.Field) logrus.Fields {
	converted := logrus.Fields{}
	for _, val := range fields {
		converted[val.Key] = val.Value
	}
	return converted
}
