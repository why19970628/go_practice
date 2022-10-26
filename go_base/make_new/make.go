package main

import "fmt"

func testMake() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 66
	a = append(a, 11, 20, 12)
	fmt.Println(a) // [66 0 0 0 0 11 20 12]

}

func main() {
	testMake()
}
