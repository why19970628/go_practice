package main

import (
	"fmt"
	"sync"
)

/*
开根号
*/
func sqrt(num float64) float64 {
	if num <= 0 {
		return 0
	}
	l, r := float64(0), num
	for r-l > 1e-3 {
		mid := l + (r-l)/2
		tmp := mid * mid
		if tmp > num {
			r = mid
		} else {
			l = mid
		}

	}
	return l
}

func routine() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	wg.Wait()
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

func findPrimes(num int) []int {
	resp := make([]int, 0)
	for i := 2; i < num; i++ {
		if isPrime(i) {
			resp = append(resp, i)
		}
	}
	return resp
}

// 判断素数
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(sqrt(2))
	resp := findPrimes(20)
	fmt.Println(resp)
}
