package main

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	book := new(Book).Builder(123, "name").SetPrice(23.4).Build()
	fmt.Println(book)
}
