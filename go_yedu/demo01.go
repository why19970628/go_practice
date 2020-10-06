package main

import "fmt"

func f() bool {
	return false
}

func main()  {
	switch f() {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}