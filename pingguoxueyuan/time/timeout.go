package main

import (
	"context"
	//"context"
	"errors"
	"fmt"
	"time"
)

func WithTimeout(dur time.Duration, fn func() error) error {

	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()

	errChan := make(chan error)
	go func() {
		err := fn()
		if err != nil {
			errChan <- errors.New("some error happen")
		} else {
			errChan <- nil
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-errChan:
		return e
	}
	return nil
}

func WithTimeout2(dur time.Duration, fn func() error) error {
	var ch = make(chan error, 1)
	go time.AfterFunc(dur, func() {
		ch <- errors.New("函数执行超时")
	})
	go func() {
		ch <- fn()
	}()
	return <-ch
}

func process() error {
	time.Sleep(time.Second * 2)
	fmt.Println("process")
	return nil
}

func main() {
	timeOut := time.Second * 3
	//err := WithTimeout2(timeOut, process)
	err := WithTimeout(timeOut, process)
	if err != nil {
		fmt.Println(err)
	}
}
