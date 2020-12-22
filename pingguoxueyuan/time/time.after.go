package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := time.After(time.Second * 3)		// 底层其实是 new Timer(d).C
	newTime := <-ch			// 阻塞3秒
	fmt.Println(newTime)
}
