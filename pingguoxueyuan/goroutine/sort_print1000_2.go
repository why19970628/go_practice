package main

import "sync"

var printCount = 5


func printAA(wg *sync.WaitGroup,chanA, chanB chan struct{})  {
	defer wg.Done()
	for i := 0; i <printCount ; i++ {
			<- chanA
			println("A")
			chanB <- struct{}{}
	}
}

func printBB(wg *sync.WaitGroup, chanB, chanC chan struct{})  {
	defer wg.Done()
	for i := 0; i <printCount ; i++ {
			<- chanB
			println("B")
			chanC <- struct{}{}
	}
}

func printCC(wg *sync.WaitGroup, chanC, chanA chan struct{})  {
	defer wg.Done()
	for i := 0; i <printCount ; i++ {
		<- chanC
		println("C")
		chanA <- struct{}{}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	chanC := make(chan struct{}, 1)
	chanA <- struct{}{}
	go printAA(&wg, chanA, chanB)
	go printBB(&wg, chanB, chanC)
	go printCC(&wg, chanC, chanA)
	wg.Wait()
}
