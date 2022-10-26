package adapter

import "testing"

var expect = "adaptee method"
var expect2 = "adaptee method2"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}

	res2 := target.Request2()
	if res2 != expect2 {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
