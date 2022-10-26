package model

import "fmt"

/***
 *  Factory Method
 */
type Factory interface {
	Create() Food
}

type MeatFactory struct {
}

func (mf MeatFactory) Create() Food {
	m := Meat{}
	return m
}

type HambergerFactory struct {
}

func (hf HambergerFactory) Create() Food {
	h := Hamberger{}
	return h
}

type Food interface {
	Eat()
}

type Meat struct {
}

type Hamberger struct {
}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}

func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

func main() {
	// Factory Method
	mf := MeatFactory{}
	mf.Create().Eat()
	hf := HambergerFactory{}
	hf.Create().Eat()
}
