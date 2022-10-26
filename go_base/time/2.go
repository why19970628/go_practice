package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("煎鱼还没到锅里...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("结束")
}
