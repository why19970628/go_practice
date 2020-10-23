package main

import (
	"fmt"
	"time"
)

func main()  {
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("%T%s\n", timer, timer)
	fmt.Println(time.Now())

	ch2 := timer.C
	fmt.Println(<-ch2)
}
