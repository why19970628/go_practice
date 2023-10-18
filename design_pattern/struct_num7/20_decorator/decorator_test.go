package decorator

import (
	"fmt"
	"testing"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()

	if got != "this is a square, color is red" {
		t.Fail()
	}
	fmt.Println(got)
}
