package main

func fib(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 0; i < n; i++ {
		c := a + b
		a, b = b, c
	}
	return c
}
