package middleware

import (
	"context"
	"net/http"

	"github.com/xMADAN05/go-common/common/dao"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API Key Required",
			})
			return
		}

		daoClient, err := dao.NewDAO("api_key_store", "us-east-2")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to connect to DynamoDB.",
			})
			return
		}
		rec, err := daoClient.GetByID(context.TODO(), apiKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid API Key.",
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
