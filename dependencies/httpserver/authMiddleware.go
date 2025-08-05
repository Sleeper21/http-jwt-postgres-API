package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
