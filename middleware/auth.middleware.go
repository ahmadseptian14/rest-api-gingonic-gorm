package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Token not valid",
		})
		return
	}

	ctx.Next()
}
