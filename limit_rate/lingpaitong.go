package main

import (
	"net/http"
	"time"
	"github.com/juju/ratelimit"
	"github.com/gin-gonic/gin"
)

// https://github.com/juju/ratelimit
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}