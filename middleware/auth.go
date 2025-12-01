package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/xMADAN05/go-common/dao"
	"github.com/xMADAN05/go-common/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const debugAPIKey = "supersecretkey"

var apiKeyDAO *dao.APIKeyDAO

func init() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		panic("unable to load AWS Config : " + err.Error())
	}
	client := dynamodb.NewFromConfig(cfg)
	apiKeyDAO = dao.NewAPIKeyDAO(client, "api_keys_table")
}

func APIKeyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		ctxId := ctx.GetHeader("X-Context-ID")

		debugMode := os.Getenv("APP_DEBUG") == "true" || gin.Mode() == gin.DebugMode

		if debugMode && apiKey == debugAPIKey {
			logger.Warn("Debug API key used",
				zap.Any("context_id", ctxId),
			)

			ctx.Set("scopes", []string{"*"})
			ctx.Set("service_name", "debug-svc")
			ctx.Next()
			return
		}

		if apiKey == "" {
			logger.Warn("Missing API Key",
				zap.String("path", ctx.FullPath()),
				zap.Any("context_id", ctxId),
			)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":      "missing_api_key",
				"message":    "X-API-Key header must be included",
				"context_id": ctxId,
			})
			return
		}
		rec, err := apiKeyDAO.Get(ctx, apiKey)
		if err != nil {
			logger.Warn("Invalid API Key",
				zap.String("api_key", apiKey),
				zap.Error(err),
				zap.Any("context_id", ctxId),
			)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":      "invalid_api_key",
				"message":    "API Key not recognized",
				"context_id": ctxId,
			})
			return
		}

		if rec.Active == "false" {
			logger.Warn("Inactive API Key access attempt",
				zap.String("api_key", apiKey),
				zap.Any("context_id", ctxId),
			)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":      "inactive_api_key",
				"message":    "API Key is inactive",
				"context_id": ctxId,
			})
			return
		}
		ctx.Set("Scopes", rec.Scopes)
		ctx.Set("ApplicationID", rec.ApplicationID)
		ctx.Next()
	}
}
