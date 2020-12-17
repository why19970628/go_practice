package main

import "fmt"

type People1 struct{}

func (p *People1) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People1) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People1
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

