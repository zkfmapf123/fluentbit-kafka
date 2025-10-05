package internal

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {

	// NewProduction은 기본적으로 stderr 로 발생
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(os.Stdout), // stdout
		zap.InfoLevel,
	)

	logger = zap.New(core, zap.AddCaller())
}

// performance -> zap.logger
// development -> zap.logger.sugar()
func NewLogger() *zap.SugaredLogger {
	sugar := logger.Sugar()
	return sugar

}
