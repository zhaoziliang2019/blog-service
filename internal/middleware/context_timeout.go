package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"time"
)

func ContextTimeout(t time.Duration) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx.Request.Context(), t)
		defer cancel()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
