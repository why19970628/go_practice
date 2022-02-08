package main

import "fmt"

func climbStairs2(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
func main() {
	fmt.Println(climbStairs2(2))
}
