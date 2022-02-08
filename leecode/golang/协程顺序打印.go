package main

import "fmt"

var index = make(chan bool, 1)
var ch1 = make(chan bool, 1)
var ch2 = make(chan bool)

func letter() {
	for i := 'A'; i <= 'Z'; i += 2 {
		<-ch1
		fmt.Println("a")
		ch2 <- true
	}
}

func number() {
	for i := 0; i < 10; i++ {
		<-ch2
		fmt.Println(i)
		ch1 <- true
	}
	index <- true

}

func main() {
	ch1 <- true
	go letter()
	go number()
	<-index
}
