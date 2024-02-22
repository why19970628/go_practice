package main

import "fmt"

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.ch
}

func main() {
	mutex := NewMutex()

	go func() {
		mutex.Lock()
		fmt.Println("Goroutine 1: Acquired lock")
		// Critical section
		mutex.Unlock()
		fmt.Println("Goroutine 1: Released lock")
	}()

	go func() {
		mutex.Lock()
		fmt.Println("Goroutine 2: Acquired lock")
		// Critical section
		mutex.Unlock()
		fmt.Println("Goroutine 2: Released lock")
	}()

	// Wait for goroutines to finish
	fmt.Scanln()
}
