package main

import "fmt"

func main() {
	//a()
	a2()
}
func a() {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) // [100 200 3]
		}
		a[k] = v + 100
	}
	fmt.Println(a) // [101 300 103]
}

func a2() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) // [100 200 3]
		}
		a[k] = v + 100
	}
	fmt.Println(a) // [101 300 103]
}
