package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Aggregator struct {
	data           map[string]int
	mu             sync.Mutex
	aggregationMap map[string]int
	aggregationDur time.Duration
	wg             sync.WaitGroup
	shutdownCh     chan struct{}
}

func NewAggregator(aggregationDur time.Duration) *Aggregator {
	return &Aggregator{
		data:           make(map[string]int),
		aggregationMap: make(map[string]int),
		aggregationDur: aggregationDur,
		shutdownCh:     make(chan struct{}),
	}
}
func (a *Aggregator) Increment(key string, value int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.aggregationMap[key]; !ok {
		a.aggregationMap[key] = value
	} else {
		a.aggregationMap[key] += value
	}
}

func (a *Aggregator) StartAggregation() {
	ticker := time.NewTicker(a.aggregationDur)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a.aggregateData()
		}
	}
}

func (a *Aggregator) aggregateData() {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 遍历聚合映射
	for key, value := range a.aggregationMap {
		// 合并聚合值
		a.data[key] += value
		if key == "user1_experience" {
			atomic.AddInt64(&a1, int64(a.data[key]))
			a.data[key] = 0
		}
		if key == "user2_experience" {
			atomic.AddInt64(&a2, int64(a.data[key]))
			a.data[key] = 0
		}
		fmt.Printf("Aggregated data for key %s: %d\n", key, a.data[key])
	}

	// 清空聚合映射
	a.aggregationMap = make(map[string]int)
}

func (a *Aggregator) StartAggregationWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		a.wg.Add(1)
		go a.aggregateDataWorker()
	}
}

func (a *Aggregator) aggregateDataWorker() {
	defer a.wg.Done()

	for {
		select {
		case <-a.shutdownCh:
			return // 退出协程
		default:
			a.aggregateData()
			// 添加适当的休眠或调整以控制聚合频率
			time.Sleep(a.aggregationDur)
		}
	}
}

func (a *Aggregator) Shutdown() {
	close(a.shutdownCh)
}

func (a *Aggregator) Wait() {
	a.wg.Wait()
}

var (
	a1 int64
	a2 int64
)

func main() {
	aggregator := NewAggregator(2 * time.Second)
	aggregator.StartAggregationWorkers(5) // 启动5个并发协程处理聚合数据

	// 模拟用户并发写入操作
	go func() {
		for i := 0; i < 100; i++ {
			aggregator.Increment("user1_experience", 1)
			aggregator.Increment("user2_experience", 2)
			time.Sleep(1 * time.Millisecond * 1)
		}
	}()
	// 捕获终止信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	quit(sigCh)

	select {
	case sig := <-sigCh:
		fmt.Printf("Received signal: %v\n", sig)

		// 执行优雅关闭
		aggregator.Shutdown()

		fmt.Println("Waiting for aggregation workers to complete...")
		aggregator.Wait()

		fmt.Println("All tasks completed")

		fmt.Println("a1", a1)
		fmt.Println("a2", a2)
		// 可在这里添加其他清理工作

		fmt.Println("Graceful shutdown completed")
	}

}

func quit(sigCh chan<- os.Signal) {
	go func() {
		time.Sleep(time.Second * 2)
		sigCh <- syscall.SIGINT
	}()
}
