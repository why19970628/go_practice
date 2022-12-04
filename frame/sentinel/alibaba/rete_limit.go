package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
)

// 流控基本字段解释
//
//Resource：资源名，即规则的作用目标。
//TokenCalculateStrategy: 当前流量控制器的Token计算策略。Direct表示直接使用字段 Threshold 作为阈值；WarmUp表示使用预热方式计算Token的阈值。
//ControlBehavior: 表示流量控制器的控制策略；Reject表示超过阈值直接拒绝，Throttling表示匀速排队。
//Threshold: 表示流控阈值；如果字段 StatIntervalInMs 是1000(也就是1秒)，那么Threshold就表示QPS，流量控制器也就会依据资源的QPS来做流控。
//RelationStrategy: 调用关系限流策略，CurrentResource表示使用当前规则的resource做流控；AssociatedResource表示使用关联的resource做流控，关联的resource在字段 RefResource 定义；
//RefResource: 关联的resource；
//WarmUpPeriodSec: 预热的时间长度，该字段仅仅对 WarmUp 的TokenCalculateStrategy生效；
//WarmUpColdFactor: 预热的因子，默认是3，该值的设置会影响预热的速度，该字段仅仅对 WarmUp 的TokenCalculateStrategy生效；
//MaxQueueingTimeMs: 匀速排队的最大等待时间，该字段仅仅对 Throttling ControlBehavior生效；
//StatIntervalInMs: 规则对应的流量控制器的独立统计结构的统计周期。如果StatIntervalInMs是1000，也就是统计QPS。

func SentinelInit() {
	//先初始化sentinel
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("初始化sentinel 异常: %v", err)
	}
	//配置限流规则

	// 设置规则qps 最大为2，超出qps的请求直接丢弃
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              2,
			StatIntervalInMs:       1000,
		},

		// 以下规则代表每 100ms 最多通过一个请求，多余的请求将会排队等待通过，若排队时队列长度大于 500ms 则直接拒绝：
		{
			Resource:               "some-test2",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Throttling, // 流控效果为匀速排队
			Threshold:              10,              // 请求的间隔控制在 1000/10=100 ms
			MaxQueueingTimeMs:      500,             // ms 最长排队等待时间
		},
	})

	if err != nil {
		log.Fatalf("加载规则失败: %v", err)
	}
}

func main() {
	SentinelInit()
	// 模仿流量请求
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			// 使用流量控制的业务逻辑必须先实例化一个 sentinel
			// 这里是对资源 some-test 进行流量访问 qps <= 2 否则就会
			e, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
			if b != nil {
				fmt.Println("限流了")
			} else {
				fmt.Println("检查通过")
				e.Exit()
			}
		}()
	}
	<-ch
}
