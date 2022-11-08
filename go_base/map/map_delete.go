package main

import "fmt"

func main() {
	var m map[string]int
	delete(m, "oh")
	fmt.Println(m)
}
