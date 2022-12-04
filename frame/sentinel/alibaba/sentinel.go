package main

import (
	"errors"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/alibaba/sentinel-golang/util"
	"log"
	"math/rand"
	"time"
)

// Sentinel 提供了监听器去监听熔断器状态机的三种状态的转换，方便用户去自定义扩展：
// 通过三个 hook 函数，用户可以很容易拿到熔断器每次状态切换的事件，以及熔断器对应的 Rule。

// Note 1: 这里需要注意的是，监听器 hook 里面携带的规则是基于 copy 的，也就是用户在监听器里面更改 Rule 不会影响到熔断器。此外这里基于拷贝是有一定性能开销的，用户要尽可能减少无效的监听器注册。
//Note 2: 熔断器监听器的注册和清除是非线程安全的，用户必须要在服务启动时配置 Sentinel 时候就注册对应的监听器，应用运行中禁止更改熔断器状态机的监听器。
type stateChangeTestListener struct {
}

// 熔断器切换到 Closed 状态时候会调用改函数, prev代表切换前的状态，rule表示当前熔断器对应的规则
func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

// 熔断器切换到 Open 状态时候会调用改函数, prev代表切换前的状态，rule表示当前熔断器对应的规则， snapshot表示触发熔断的值
func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %d, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

// 熔断器切换到 HalfOpen 状态时候会调用改函数, prev代表切换前的状态，rule表示当前熔断器对应的规则
func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func sentinelInit() {
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	// 模拟10s 内错误数量达到10的时候熔断
	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=5s, recoveryTimeout=3s, maxErrorCount=50

		// 模拟10s 内错误数量达到10的时候熔断
		{
			Resource: "a",
			// 设置通断策略为错误数
			Strategy: circuitbreaker.ErrorCount,
			// 熔断触发后持续的时间（单位为 ms）
			RetryTimeoutMs: 1000,
			// 静默数量，对资源的访问小于静默数，熔断器处于静默状态
			MinRequestAmount: 3,
			// 熔断器的统计周期，单位是毫秒， 一般情况下设置10秒左右都OK
			StatIntervalMs: 10000,
			// 熔断器的统计周期内，统计滑动窗口的桶数，默认为 1
			// 随着桶数的增加，统计数据会更精确，但内存消耗也会增加
			// 以下必须为真- " StatIntervalMs % StatSlidingWindowBucketCount == 0 "，
			// 否则 StatSlidingWindowBucketCount 将被替换为1。
			StatSlidingWindowBucketCount: 10,
			// 错误数的阈值
			Threshold: 10,
		},

		// 错误比率熔断
		{
			Resource: "b",
			// 设置通断策略为错误比例
			Strategy: circuitbreaker.ErrorRatio,
			// 熔断触发后持续的时间（单位为 ms）
			RetryTimeoutMs: 3000,
			// 静默数量，对资源的访问小于静默数，熔断器处于静默状态
			MinRequestAmount: 10,
			// 熔断器的统计周期，单位是毫秒， 一般情况下设置10秒左右都OK
			StatIntervalMs: 10000,
			// 熔断器的统计周期内，统计滑动窗口的桶数，默认为 1
			// 随着桶数的增加，统计数据会更精确，但内存消耗也会增加
			// 以下必须为真- " StatIntervalMs % StatSlidingWindowBucketCount == 0 "，
			// 否则 StatSlidingWindowBucketCount 将被替换为1。
			StatSlidingWindowBucketCount: 10,
			// 错误比例的阈值 20.1%
			Threshold: 0.201,
		},

		// 慢响应比率熔断
		{
			Resource: "c",
			// 设置通断策略为慢调用
			Strategy: circuitbreaker.SlowRequestRatio,
			// 熔断触发后持续的时间（单位为 ms）
			RetryTimeoutMs: 3000,
			// 静默数量，对资源的访问小于静默数，熔断器处于静默状态
			MinRequestAmount: 10,
			// 熔断器的统计周期，单位是毫秒， 一般情况下设置10秒左右都OK
			StatIntervalMs: 5000,
			// 慢调用响应时间阈值，RT大于该值的请求判断为慢响应
			MaxAllowedRtMs: 50,
			// 慢调用比例的阈值(小数表示，比如0.1表示10%)
			Threshold: 0.5,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	logging.Info("[CircuitBreaker ErrorCount] Sentinel Go circuit breaking demo is running. You may see the pass/block metric in the metric log.")

}

func main() {
	sentinelInit()
	ch := make(chan struct{})
	//abc()
	d()

	<-ch
}

func abc() {
	go func() {
		for {
			e, b := sentinel.Entry("a")
			if b != nil {
				// g1 blocked
				fmt.Println("g1熔断")
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				fmt.Println("g1通过")
				// 模拟访问资源错误逻辑
				if rand.Uint64()%20 > 9 {
					fmt.Println("访问资源失败")
					// Record current invocation as error.
					// 当对资源访问失败的时候必须向sentinel 发送一个错误追踪
					sentinel.TraceError(e, errors.New("biz error"))
				}
				// g1 passed
				time.Sleep(time.Duration(rand.Uint64()%80+10) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	go func() {
		for {
			e, b := sentinel.Entry("a")
			if b != nil {
				fmt.Println("g2熔断")
				// g2 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				fmt.Println("g2通过")
				// g2 passed
				time.Sleep(time.Duration(rand.Uint64()%80) * time.Millisecond)
				e.Exit()
			}
		}
	}()
}

func d() {
	go func() {
		for {
			e, b := sentinel.Entry("c")
			if b != nil {
				// g1 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				if rand.Uint64()%20 > 9 {
					// Record current invocation as error.
					sentinel.TraceError(e, errors.New("biz error"))
				}
				// g1 passed
				// 这里sleep模拟随机慢响应，随机值在10——90ms内
				time.Sleep(time.Duration(rand.Uint64()%80+10) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	go func() {
		for {
			e, b := sentinel.Entry("c")
			if b != nil {
				// g2 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				// g2 passed
				time.Sleep(time.Duration(rand.Uint64()%80) * time.Millisecond)
				e.Exit()
			}
		}
	}()
}
