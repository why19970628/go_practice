package main

import (
	"errors"
	"fmt"
	"time"
)

const defaultTimeout time.Duration = time.Millisecond * 100

// 利用一个超时控制，实现一个函数超时
func AsyncCall(t time.Duration, f func() error) error {
	done := make(chan struct{}, 1)
	errChan := make(chan error, 1)

	go func() {
		err := f()
		if err != nil {
			errChan <- err
		} else {
			done <- struct{}{}
		}
	}()

	select {
	case <-done:
		fmt.Println("call successfully!!!")
		return nil
	case e := <-errChan:
		return e
	case <-time.After(defaultTimeout):
		return errors.New("timeout")
	}
}

func main() {
	err := AsyncCall(time.Millisecond*50, func() error {
		time.Sleep(time.Millisecond * 100)
		return nil
	})
	if err != nil {
		fmt.Println("err::", err)
	}
}
