package config

import (
	"fmt"
	"testing"
)

func TestOperator(t *testing.T) {
	// iterator

	// ============== type1 ==================
	aggregate1 := NewConcreteAggregate()

	aggregate1.AddElement("youzg")
	aggregate1.AddElement("is")
	aggregate1.AddElement("a")
	aggregate1.AddElement("good")
	aggregate1.AddElement("man")
	aggregate1.AddElement("!")

	iterator1 := aggregate1.Iterator()

	iterator1.Remove()

	for iterator1.HashNext() {
		fmt.Println(iterator1.Next())
	}
}
