package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Num      int
	Result   int
	Finished bool
}

func Worker(ctx context.Context, id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				// 任务通道已关闭，退出worker
				fmt.Printf("Worker %d received exit signal\n", id)
				return
			}

			fmt.Printf("Worker %d started task %d\n", id, task.ID)
			task.Result = task.Num * task.Num // 假设任务是计算输入数的平方
			task.Finished = true
			fmt.Printf("Worker %d finished task %d\n", id, task.ID)
		case <-time.After(5 * time.Second):
			// 超时5秒，退出worker
			fmt.Printf("Worker %d timeout, exiting\n", id)
			return
		case <-ctx.Done():
			// 收到退出信号，退出worker
			fmt.Printf("Worker %d received exit signal\n", id)
			return
		}
	}
}

func Master(numWorkers int, numTasks int) {
	tasks := make(chan Task, numTasks)                      // 创建任务通道
	ctx, cancel := context.WithCancel(context.Background()) // 创建带取消功能的context
	var wg sync.WaitGroup

	// 启动worker协程
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(ctx, i, tasks, &wg)
	}

	// 分发任务给worker
	for i := 1; i <= numTasks; i++ {
		task := Task{
			ID:       i,
			Num:      i,
			Result:   0,
			Finished: false,
		}
		tasks <- task
	}

	// 关闭任务通道，等待所有worker完成任务
	close(tasks)
	wg.Wait()

	// 发送退出信号给所有worker
	cancel()

	fmt.Println("Master finished")
}

func main() {
	numWorkers := 3
	numTasks := 1000

	go func() {
		time.Sleep(5 * time.Second) // 模拟Master运行5秒后退出
		fmt.Println("Master timeout, sending exit signal")
		// 发送退出信号给Master
		// 在这里可以进行一些清理工作
	}()

	Master(numWorkers, numTasks)
}
