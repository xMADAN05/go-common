package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

func Init(level, env string) error {
	var err error

	once.Do(func() {
		if level == "" {
			level = "info"
		}
		if env == "" {
			env = "dev"
		}

		lvl, err2 := zapcore.ParseLevel(level)

		if err2 != nil {
			err = err2
			return
		}

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		encoding := "json"
		if env == "dev" {
			encoding = "console"
			encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		cfg := zap.Config{
			Level:            zap.NewAtomicLevelAt(lvl),
			Encoding:         encoding,
			EncoderConfig:    encoderCfg,
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		log, err = cfg.Build()
	})

	return err
}

// USECASE: Containerized apps; 12-factor configuration.
func InitFromEnv() error {
	return Init(os.Getenv("LOG_LEVEL"), os.Getenv("ENV"))
}

// DOES: Returns the global logger. Automatically initializes with defaults if not already done.
// USECASE: Internal helper + fallback for missed initialization.
func L() *zap.Logger {
	if log == nil {
		_ = Init("", "")
	}
	return log
}

func Debug(msg string, f ...zap.Field) { L().Debug(msg, f...) }
func Info(msg string, f ...zap.Field)  { L().Info(msg, f...) }
func Warn(msg string, f ...zap.Field)  { L().Warn(msg, f...) }
func Error(msg string, f ...zap.Field) { L().Error(msg, f...) }
func Fatal(msg string, f ...zap.Field) { L().Fatal(msg, f...) }

func With(f ...zap.Field) *zap.Logger { return L().With(f...) }

// DOES: Flushes buffered logs to stdout/stderr.
// USECASE: Required for AWS Lambda, CLI tools, and graceful shutdowns.
func Sync() error {
	if log != nil {
		return log.Sync()
	}
	return nil
}
