package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Item represents an item in the priority queue.
type Item struct {
	ID    string
	Count int
}

// PriorityQueue implements a min-heap for Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Count < pq[j].Count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

const (
	windowSize       = 10 // 滑动窗口大小，根据需要调整
	topK             = 5  // Top-K数量，根据需要调整
	callbackInterval = 5  // 回调间隔，根据需要调整（秒）
	userRequestRate  = 10 // 用户请求速率（每秒请求数），根据需要调整
)

var (
	hotspots      = make(map[string]int)
	hotspotsMutex = &sync.Mutex{}
	window        = make([]map[string]int, windowSize)
	windowMutex   = &sync.Mutex{}
)

/*

这个示例代码包括了一个滑动窗口来跟踪用户请求的统计数据，以实现秒级感知。
滑动窗口的大小是windowSize，用户请求的速率是userRequestRate，你可以根据需要调整这些值。
当定时回调触发时，它将使用滑动窗口+LFU算法计算Top-K，并将热点数据ID回调给业务。
在实际应用中，你可以根据需要进一步优化和扩展这个基本框架。

*/
func main() {
	// 初始化滑动窗口
	for i := 0; i < windowSize; i++ {
		window[i] = make(map[string]int)
	}

	// 启动定时回调
	go periodicCallback()

	// 模拟用户请求，调用计数API
	for {
		for i := 0; i < userRequestRate; i++ {
			userRequest("data_id")
		}
		time.Sleep(time.Second)
	}

	// 阻塞主线程
	select {}
}

// 用户请求处理，调用计数API
func userRequest(dataID string) {
	// 实际计数逻辑
	count(dataID)
}

// 计数逻辑，更新热点数据统计
func count(dataID string) {
	hotspotsMutex.Lock()
	defer hotspotsMutex.Unlock()

	windowMutex.Lock()
	defer windowMutex.Unlock()

	hotspots[dataID]++
	window[0][dataID]++
}

// 定时回调，将统计到的热点数据ID回调给业务
func periodicCallback() {
	for {
		time.Sleep(callbackInterval * time.Second)

		hotspotsMutex.Lock()
		windowMutex.Lock()

		// 合并窗口数据到热点数据
		for i := 0; i < windowSize-1; i++ {
			for id, count := range window[i] {
				hotspots[id] += count
			}
			window[i] = make(map[string]int)
			//copy(window[i], window[i+1])
		}
		window[windowSize-1] = make(map[string]int)

		// 计算Top-K
		if len(hotspots) > 0 {
			// 将统计到的热点数据ID按照LFU算法计算Top-K
			topKItems := calculateTopK()
			// 清空热点数据
			hotspots = make(map[string]int)
			hotspotsMutex.Unlock()
			windowMutex.Unlock()

			// 回调业务，传递Top-K数据ID
			callbackToBusiness(topKItems)
		} else {
			hotspotsMutex.Unlock()
			windowMutex.Unlock()
		}
	}
}

// 使用LFU算法计算Top-K热点数据
func calculateTopK() []*Item {
	pq := make(PriorityQueue, 0, topK)
	for id, count := range hotspots {
		if len(pq) < topK {
			heap.Push(&pq, &Item{ID: id, Count: count})
		} else {
			if count > pq[0].Count {
				heap.Pop(&pq)
				heap.Push(&pq, &Item{ID: id, Count: count})
			}
		}
	}

	// 将队列中的数据按照Count从大到小排序
	topKItems := make([]*Item, topK)
	for i := topK - 1; i >= 0; i-- {
		topKItems[i] = heap.Pop(&pq).(*Item)
	}

	return topKItems
}

// 模拟回调业务，将Top-K数据ID传递给业务
func callbackToBusiness(topKItems []*Item) {
	fmt.Println("Callback to business with Top-K data:")
	for _, item := range topKItems {
		fmt.Printf("ID: %s, Count: %d\n", item.ID, item.Count)
	}
}
