package internal

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	l *zap.SugaredLogger
}

type ILogger interface {
	InfoLogger(funcname string, msgs map[string]any)
	ErrorLogger(funcname string, msgs map[string]any)
}

// performance -> zap.logger
// development -> zap.logger.sugar()
func NewLogger() *Logger {

	// NewProduction은 기본적으로 stderr 로 발생
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(os.Stdout), // stdout
		zap.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	return &Logger{
		l: logger.Sugar(),
	}
}

func (l *Logger) logFormat(funcname string, msgs map[string]any, logLevel string) {

	// SugaredLogger의 Infow, Warnw 등은 key-value 쌍을 받습니다
	fields := make([]interface{}, 0, len(msgs)*2)
	for k, v := range msgs {
		fields = append(fields, k, v)
	}

	switch logLevel {
	case "info":
		l.l.Infow(funcname, fields...)
	case "warn":
		l.l.Warnw(funcname, fields...)
	case "debug":
		l.l.Debugw(funcname, fields...)
	case "error":
		l.l.Errorw(funcname, fields...)
	}
}

func (l *Logger) InfoLogger(funcname string, msgs map[string]any) {
	l.logFormat(funcname, msgs, "info")
}

func (l *Logger) ErrorLogger(funcname string, msgs map[string]any) {
	l.logFormat(funcname, msgs, "warn")
}

func (l *Logger) Sync() {
	l.l.Sync()
}
