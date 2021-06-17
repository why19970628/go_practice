package main

import (
	"fmt"
	"sync"
)

//控制循环次数
var count = 1000

func main() {
	wg := sync.WaitGroup{}
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	chanC := make(chan struct{}, 1)
	chanA <- struct{}{}
	wg.Add(3)
	go printA(&wg, chanA, chanB)
	go printB(&wg, chanB, chanC)
	go printC(&wg, chanC, chanA)
	wg.Wait()
}
func printA(wg *sync.WaitGroup, chanA, chanB chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanA
		fmt.Println("A")
		chanB <- struct{}{}
	}
}
func printB(wg *sync.WaitGroup, chanB, chanC chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanB
		fmt.Println("B")
		chanC <- struct{}{}
	}
}
func printC(wg *sync.WaitGroup, chanC, chanA chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanC
		fmt.Println("C")
		chanA <- struct{}{}
	}
}
