package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// https://www.liwenzhou.com/posts/Go/ratelimit/#autoid-0-2-1

// 创建令牌桶的方法：

	// 创建指定填充速率和容量大小的令牌桶
	//func NewBucket(fillInterval time.Duration, capacity int64) *Bucket

	// 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
	//func NewBucketWithQuantum(fillInterval time.Duration, capacity, quantum int64) *Bucket

	// 创建填充速度为指定速率和容量大小的令牌桶
	// NewBucketWithRate(0.1, 200) 表示每秒填充20个令牌
	//func NewBucketWithRate(rate float64, capacity int64) *Bucket



//取出令牌的方法如下：


	// 取token（非阻塞）
	//func (tb *Bucket) Take(count int64) time.Duration
	//func (tb *Bucket) TakeAvailable(count int64) int64
	//// 最多等maxWait时间取token
	//func (tb *Bucket) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool)
	//
	//// 取token（非阻塞）
	//func (tb *Bucket) Wait(count int64)
	//func (tb *Bucket) Wa
	//
	//itMaxDuration(count int64, maxWait time.Duration) bool


// 对于该限流中间件的注册位置，我们可以按照不同的限流策略将其注册到不同的位置，例如：

// 如果要对全站限流就可以注册成全局的中间件。
// 如果是某一组路由需要限流，那么就只需将该限流中间件注册到对应的路由组即可。


func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) > 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}