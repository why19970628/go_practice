package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sync.Cond实现了一个条件变量，用于等待一个或一组goroutines满足条件后唤醒的场景。每个Cond关联一个Locker通常是一个*Mutex或RWMutex`根据需求初始化不同的锁。
// https://www.jianshu.com/p/1d3d315fbd7a

type Fifo struct {
	queue []int
	m     *sync.Mutex
	cond  *sync.Cond
}

type Queue interface {
	Offer(num int) error
	Pop() int
}

func (f *Fifo) Offer(num int) error {
	f.m.Lock()
	defer f.m.Unlock()
	f.queue = append(f.queue, num)
	// 唤醒全部等待的goroutine
	f.cond.Broadcast()
	return nil
}

func (f *Fifo) Pop() int {
	f.m.Lock()
	defer f.m.Unlock()
	for {
		if len(f.queue) == 0 {
			// 将当前goroutine插入到notifyList链表中
			// 调用gopark将当前goroutine挂起
			f.cond.Wait()
		}
		v := f.queue[0]
		f.queue = f.queue[1:]
		return v
	}
}

func main() {
	m := &sync.Mutex{}
	f := &Fifo{
		queue: []int{},
		m:     m,
		cond:  sync.NewCond(m),
	}

	count := 0
	// offer
	go func() {
		for {
			if count > 10 {
				break
			}
			_ = f.Offer(rand.Int())
			count++
		}
	}()

	time.Sleep(time.Second)

	// pop
	go func() {
		for {
			v := f.Pop()
			fmt.Println("g1:", v)
		}
	}()

	go func() {
		for {
			v := f.Pop()
			fmt.Println("g2:", v)
		}
	}()

	time.Sleep(time.Second * 2)

	// offer
	//go func() {
	//	for  {
	//		f.Offer(rand.Int())
	//		fmt.Println(f.queue)
	//	}
	//}()

	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, os.Interrupt)
	//<-ch

}
