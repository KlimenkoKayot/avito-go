package domain

type LogField interface{}

type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	With(fields ...LogField) Logger
}
