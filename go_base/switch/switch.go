package main

import "fmt"

func testIf() {
	var a int = 10
	switch a {
	case 1:
		fmt.Printf("a=1\n")

	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	default:
		fmt.Println("on num")

	}
}
func main() {
	testIf()
}
