package main

import "fmt"

type People2 struct {
	Name string
}

func (p *People2) String() {
	fmt.Printf("print: %v", p)
}

func main() {
	p := &People2{
		//Name: "123",
	}
	p.String()
}
