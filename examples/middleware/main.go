package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xMADAN05/go-common/common/config"
	"github.com/xMADAN05/go-common/common/logger"
	"github.com/xMADAN05/go-common/common/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()
	gin.SetMode(cfg.GinMode)
	r := gin.Default()
	r.Use(middleware.ContextIDMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctxID, _ := ctx.Get("context_id")
		logger.Info("request recieved", zap.String("context_id", ctxID.(string)))
		ctx.JSON(http.StatusOK, gin.H{
			"context_id": ctxID,
		})
	})

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)
	res, err := http.Get(fmt.Sprintf("http://localhost%s/", addr))
	if err != nil {
		logger.Error("Failed to ping self", zap.Error(err))
	} else {
		defer res.Body.Close()
		logger.Info("Pinged self", zap.Int("status", res.StatusCode))
	}
}
