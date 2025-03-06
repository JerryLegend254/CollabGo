package logger

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() Logger {
	zapSLogger := zap.Must(zap.NewProduction()).Sugar()
	return Logger{zapSLogger}
}
