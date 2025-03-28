package logger

type Field struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

func String(msg, value string) Field {
	return Field{msg, value}
}

func Int(msg string, value int) Field {
	return Field{msg, value}
}

func Error(value error) Field {
	return Field{Value: value}
}
