package middleware

import (
	"mygram-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := helper.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", token)
		ctx.Next()
	}
}
