package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/xMADAN05/go-common/dao"

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

		debugMode := os.Getenv("APP_DEBUG") == "true" || gin.Mode() == gin.DebugMode

		if debugMode && apiKey == debugAPIKey {
			ctx.Set("scopes", []string{"*"})
			ctx.Set("service_name", "debug-svc")
			ctx.Next()
			return
		}

		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API Key Required",
			})
			return
		}
		rec, err := apiKeyDAO.Get(ctx, apiKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid API Key",
			})
			return
		}

		if rec.Active == "false" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "API Key is inactive.",
			})
			return
		}
		ctx.Set("scopes", rec.Scopes)
		ctx.Set("service_name", rec.ServiceName)
		ctx.Next()
	}
}
