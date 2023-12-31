package log

import (
	"sync"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once sync.Once = sync.Once{}
	logger *zap.Logger
)

func InitLogger(filePath string) (err error) {
	once.Do(func() {
		cfg := zap.NewProductionConfig()
		if len(filePath) > 0 {
			cfg.OutputPaths = append(cfg.OutputPaths, filePath)
		}
		logger,err = cfg.Build()
	})
	return
}

type ILogger interface {
	Debug(string, ...zapcore.Field)
	Info(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	Warn(string, ...zapcore.Field)
}

func GetLogger() ILogger {
	return logger
}