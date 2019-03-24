package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleErrors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": ctx.Errors.Last().Error(),
		})
	}
}
