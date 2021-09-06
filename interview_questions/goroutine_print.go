package main

//func main() {
//	c := make(chan int)
//	run := func() {
//		for {
//			i, ok := <-c
//			if !ok {
//				break
//			}
//			if i > 10 {
//				close(c)
//				break
//			}
//			fmt.Println(i)
//			c <- i + 1
//		}
//	}
//	go run()
//	c <- 1
//	run()
//}

func main() {
	ch := make(chan struct{})
	run := func(i, n int) {
		for ; i < n; i += 2 {
			<-ch
			println(i)
			ch <- struct{}{}
		}
		if _, ok := <-ch; ok {
			close(ch)
		}
	}

	go run(1, 10)
	ch <- struct{}{}
	run(2, 11)
}
