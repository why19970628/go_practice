package proxy

import (
	"log"
	"testing"
)

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

func TestProxy023(t *testing.T) {
	log.Printf("1---UserGroup: %+v\n", 1)
}
