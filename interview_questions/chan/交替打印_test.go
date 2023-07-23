package _chan

import (
	"fmt"
	"sync"
	"testing"
)

// 交替打印
func TestPrintV1(t *testing.T) {
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	//chanExist := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go PrintA(chanA, chanB, wg)
	go PrintB(chanA, chanB, wg)
	chanA <- struct{}{}
	//<-chanExist
	wg.Wait()
}

func PrintA(chanA, chanB chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		<-chanA
		fmt.Println("a", i)
		chanB <- struct{}{}
	}
}

func PrintB(chanA, chanB chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		<-chanB
		fmt.Println("b", i)
		if i != 49 {
			chanA <- struct{}{}
		}
	}
}

func TestPrintV2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan struct{})
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	wg.Wait()
}

func TestPrintV3(t *testing.T) {
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanExist := make(chan struct{})

	go func() {
		for i := 0; i < 50; i++ {
			<-chanA
			fmt.Println("a", i)
			chanB <- struct{}{}
		}
	}()
	go func() {
		defer func() { close(chanExist) }()
		for i := 0; i < 50; i++ {
			<-chanB
			fmt.Println("b", i)
			if i != 49 {
				chanA <- struct{}{}
			}
		}
	}()

	chanA <- struct{}{}
	<-chanExist
}
