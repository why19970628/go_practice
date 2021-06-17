package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 打乱数组顺序

func init() {
	rand.Seed(time.Now().Unix())
}

func Shuffle(arr []int) {
	for len(arr) > 0 {
		n := len(arr)
		randIndex := rand.Intn(n)
		fmt.Println("randIndex:", randIndex)
		arr[n-1], arr[randIndex] = arr[randIndex], arr[n-1]
		arr = arr[:n-1]
		fmt.Println(arr)
	}
}
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Println(arr)
	Shuffle(arr)
	fmt.Println(arr)
}
