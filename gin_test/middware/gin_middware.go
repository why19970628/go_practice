package middware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// docs: https://www.liwenzhou.com/posts/Go/ratelimit/
// 创建一个容量为5的桶,并且每2秒往桶里放一枚令牌
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		// 返回的是此次取出的token数量，如果返回0表示此时没有可用的token
		// https://github.com/juju/ratelimit/blob/f60b32039441cd828005f82f3a54aafd00bc9882/ratelimit.go#L221

		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.New()
	fillInterval := 2 * time.Second
	//prometheus 监控
	p := ginprometheus.NewPrometheus("go_practice")
	p.Use(r)
	r.GET("/", RateLimitMiddleware(fillInterval, 5), func(c *gin.Context) {
		c.String(http.StatusOK, "this is main")
	})
	r.Run(":8080")
}
