package main

import "runtime"

func main() {
	//defer println("bye") //不会打印
	//
	//os.Exit(1)


	c := make(chan  int)

	go func() {
		defer close(c)
		defer println("bye") //会打印
		runtime.Goexit()
	}()

	<-c
}
