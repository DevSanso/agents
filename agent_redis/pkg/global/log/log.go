package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"errors"
)

var (
	once   sync.Once = sync.Once{}
	logger *zap.Logger
)

type LogLevel string

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel LogLevel = "DEBUG"
	// InfoLevel is the default logging priority.
	InfoLevel = "INFO"
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = "WARN"
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = "ERROR"
)

var (
	ErrorInvalidLogLevel = errors.New("invalid log level")
)

func convertLevel(level LogLevel) (zap.AtomicLevel, error) {
	var ret zap.AtomicLevel
	var err error = nil
	switch level {
		case "DEBUG":
			ret = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "INFO":
			ret = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "WARN":
			ret = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "ERROR":
			ret = zap.NewAtomicLevelAt(zap.ErrorLevel)
		default:
			err = ErrorInvalidLogLevel

	}
	return ret, err
}

func InitLogger(filePath string, level LogLevel) (err error) {
	once.Do(func() {
		cfg := zap.NewProductionConfig()
		if len(filePath) > 0 {
			f, fErr := os.OpenFile(filePath, os.O_CREATE, 0666)
			if fErr != nil {
				err = fErr
				return
			}
			f.Close()

			cfg.OutputPaths = append(cfg.OutputPaths, filePath)
		}
		cfg.Level,err = convertLevel(level)
		if err != nil {
			return	
		}
		cfg.DisableStacktrace = true
		
		if cfg.Level.Level() >= zap.ErrorLevel {
			cfg.DisableCaller = true
		}
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = cfg.Build()
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
