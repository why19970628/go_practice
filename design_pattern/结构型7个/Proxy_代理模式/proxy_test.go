package proxy

import "testing"

func TestProxy1(t *testing.T) {

	var p Proxy1
	p = newNow1("now")
	p.Say()

}

func TestProxy0(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}
