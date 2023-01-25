package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoziliang2019/blog-service/pkg/app"
	"github.com/zhaoziliang2019/blog-service/pkg/errcode"
	"github.com/zhaoziliang2019/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(context *gin.Context) {
		key := l.Key(context)
		if bucket, ok := l.GetBuket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(context)
				response.ToErrorResponse(errcode.TooManyRequests)
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
