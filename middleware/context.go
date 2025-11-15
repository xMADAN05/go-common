package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const ContextKeyRequestID = "context_id"

func ContextIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contextID := ctx.GetHeader("X-Context-ID")
		if contextID == "" {
			contextID = uuid.New().String()
		}
		// Set in Gin context
		ctx.Set(ContextKeyRequestID, contextID)
		// Add to response header
		ctx.Writer.Header().Set("X-Context-ID", contextID)
		ctx.Next()
	}

}
