package _func

import (
	"go_practice/lib/helper"
	"strconv"
	"testing"
)

func TestRunMultiFunc(t *testing.T) {
	var (
		a []int
		b []string
		c = make(map[int]int)
	)

	aFn := func() {
		for i := 0; i < 10; i++ {
			a = append(a, i)
		}
	}

	bFn := func() {
		for i := 0; i < 10; i++ {
			b = append(b, strconv.Itoa(i))
		}
	}

	cFn := func() {
		for i := 0; i < 10; i++ {
			c[i] = i
		}
	}

	helper.RunMultiFunc([]func(){
		aFn,
		bFn,
		cFn,
	})
	t.Log(a)
	t.Log(b)
	t.Log(c)
}
