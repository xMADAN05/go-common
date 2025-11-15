package main

import (
	"time"

	"github.com/xMADAN05/go-common/common/logger"

	"go.uber.org/zap"
)

func main() {
	if err := logger.InitFromEnv(); err != nil {
		panic(err)
	}

	logger.Debug("This is a DEBUG message", zap.String("steps", "init"))
	logger.Info("Starting Aplication", zap.Time("time", time.Now()))
	logger.Warn("This is a Warning", zap.Int("retry", 1))
	logger.Error("Error")

	reqLogger := logger.With(
		zap.String("request_id", "b627810b-716b-486f-8ec1-f94091620620"),
		zap.String("user_id", "1000"),
	)

	reqLogger.Info("Request recieved.")

}
