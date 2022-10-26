package main

import "fmt"

type HandlersChain struct {
	A string
}

func main() {
	finalSize := 10
	mergedHandlers := make([]int, finalSize)
	fmt.Println(mergedHandlers)

	slice1 := []int{1, 2, 3, 4, 5}
	copy(mergedHandlers, slice1)
	fmt.Println(mergedHandlers)
	fmt.Println("----")

	slice2 := []int{5, 4, 3}
	fmt.Println(len(slice1), mergedHandlers[len(slice1):])
	i := copy(mergedHandlers[len(slice1):], slice2)
	fmt.Println(i, mergedHandlers)
}
