package main

import (
	"context"
	"time"
)

func FuncTimeOut(ctx context.Context, timeOut time.Duration, fn func() error) error {
	newCtx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	errChan := make(chan error)
	go func() {
		err := fn()
		errChan <- err
	}()

	for {
		select {
		case <-newCtx.Done():
			return newCtx.Err()
		case err := <-errChan:
			return err
		}
	}

}
