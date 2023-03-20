package main

import (
	"context"
	"fmt"
	"time"
)

func AsyncCall(ctx context.Context, t time.Duration, f func() error) error {

	newCtx, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	errChan := make(chan error, 1)
	panicChan := make(chan interface{}, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		errChan <- f()
	}()

	select {
	case <-newCtx.Done():
		return newCtx.Err()
	case e := <-errChan:
		return e
	case p := <-panicChan:
		panic(p)
	}

}
func main() {
	err := AsyncCall(context.Background(),
		time.Millisecond*50, func() error {
			time.Sleep(time.Millisecond * 10)
			return nil
		})
	if err != nil {
		fmt.Println("err::", err)
	}
}
